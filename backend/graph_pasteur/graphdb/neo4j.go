package graphdb

import (
	"errors"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Neo4jDB encapsulates the Neo4j database connection and provides methods for interacting with the graph database.
type Neo4jDB struct {
	driver neo4j.Driver
}

// NewNeo4jDB creates a new Neo4jDB instance with the given connection URI, username, and password.
func NewNeo4jDB(uri, username, password string) (*Neo4jDB, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	return &Neo4jDB{driver: driver}, nil
}

// Close closes the Neo4j database connection.
func (db *Neo4jDB) Close() error {
	if db.driver != nil {
		return db.driver.Close()
	}
	return nil
}

// CreateNode creates a new node with the given label and properties.
func (db *Neo4jDB) CreateNode(label string, properties map[string]interface{}) (int64, error) {
	session := db.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"CREATE (n:"+label+") SET n = $props RETURN id(n)",
			map[string]interface{}{
				"props": properties,
			},
		)
		if err != nil {
			return nil, err
		}

		summary, err := result.Consume()
		if err != nil {
			return nil, err
		}

		if summary.Next() {
			idValue, ok := summary.Record().Get("id(n)")
			if !ok {
				return nil, errors.New("Failed to retrieve node ID")
			}
			return idValue.(int64), nil
		}

		return nil, errors.New("No node ID returned")
	})

	if err != nil {
		return 0, err
	}

	return result.(int64), nil
}

// CreateRelationship creates a relationship between two nodes with the given relationship type and properties.
func (db *Neo4jDB) CreateRelationship(startNodeID, endNodeID int64, relType string, properties map[string]interface{}) (int64, error) {
	session := db.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (startNode) WHERE id(startNode) = $startNodeID "+
				"MATCH (endNode) WHERE id(endNode) = $endNodeID "+
				"CREATE (startNode)-[r:"+relType+" $props]->(endNode) RETURN id(r)",
			map[string]interface{}{
				"startNodeID": startNodeID,
				"endNodeID":   endNodeID,
				"props":       properties,
			},
		)
		if err != nil {
			return nil, err
		}

		summary, err := result.Consume()
		if err != nil {
			return nil, err
		}

		if summary.Next() {
			idValue, ok := summary.Record().Get("id(r)")
			if !ok {
				return nil, errors.New("Failed to retrieve relationship ID")
			}
			return idValue.(int64), nil
		}

		return nil, errors.New("No relationship ID returned")
	})

	if err != nil {
		return 0, err
	}

	return result.(int64), nil
}

// MatchNodes performs a graph query to match nodes with the given label and properties.
func (db *Neo4jDB) MatchNodes(label string, properties map[string]interface{}) ([]int64, error) {
	session := db.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (n:" + label + ")"

		params := make(map[string]interface{})
		if len(properties) > 0 {
			query += " WHERE "
			i := 0
			for key, value := range properties {
				query += "n." + key + " = $" + key
				params[key] = value
				i++
				if i < len(properties) {
					query += " AND "
				}
			}
		}

		query += " RETURN id(n)"

		result, err := tx.Run(query, params)
		if err != nil {
			return nil, err
		}

		var nodeIDs []int64
		for result.Next() {
			idValue, ok := result.Record().Get("id(n)")
			if !ok {
				return nil, errors.New("Failed to retrieve node ID")
			}
			nodeIDs = append(nodeIDs, idValue.(int64))
		}

		return nodeIDs, nil
	})

	if err != nil {
		return nil, err
	}

	return result.([]int64), nil
}
