package graphdb

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4j struct {
	Driver neo4j.Driver
}

func NewNeo4j(uri, username, password string) (*Neo4j, error) {
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""), configForNeo4j40)
	if err != nil {
		return nil, fmt.Errorf("failed to create Neo4j driver: %w", err)
	}

	return &Neo4j{
		Driver: driver,
	}, nil
}

func (n *Neo4j) Close() {
	n.Driver.Close()
}

func (n *Neo4j) CreateNode(label string, properties map[string]interface{}) (neo4j.Node, error) {
	session := n.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := fmt.Sprintf("CREATE (n:%s) SET n = $props RETURN n", label)
		params := map[string]interface{}{
			"props": properties,
		}
		return tx.Run(query, params)
	})

	if err != nil {
		log.Printf("Failed to create node in Neo4j: %v", err)
		return nil, err
	}

	record, err := result.Single()
	if err != nil {
		log.Printf("Failed to get node from result: %v", err)
		return nil, err
	}

	node, ok := record.GetByIndex(0).(neo4j.Node)
	if !ok {
		return nil, fmt.Errorf("failed to convert result to node")
	}

	return node, nil
}

func (n *Neo4j) CreateRelationship(startNode, endNode neo4j.Node, relType string, properties map[string]interface{}) (neo4j.Relationship, error) {
	session := n.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := fmt.Sprintf("MATCH (start), (end) WHERE id(start) = $startId AND id(end) = $endId CREATE (start)-[r:%s]->(end) SET r = $props RETURN r", relType)
		params := map[string]interface{}{
			"startId": startNode.Id(),
			"endId":   endNode.Id(),
			"props":   properties,
		}
		return tx.Run(query, params)
	})

	if err != nil {
		log.Printf("Failed to create relationship in Neo4j: %v", err)
		return nil, err
	}

	record, err := result.Single()
	if err != nil {
		log.Printf("Failed to get relationship from result: %v", err)
		return nil, err
	}

	relationship, ok := record.GetByIndex(0).(neo4j.Relationship)
	if !ok {
		return nil, fmt.Errorf("failed to convert result to relationship")
	}

	return relationship, nil
}

func (n *Neo4j) MatchNodes(label string, properties map[string]interface{}) ([]neo4j.Node, error) {
	session := n.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := fmt.Sprintf("MATCH (n:%s) WHERE n = $props RETURN n", label)
		params := map[string]interface{}{
			"props": properties,
		}
		return tx.Run(query, params)
	})

	if err != nil {
		log.Printf("Failed to execute MATCH query in Neo4j: %v", err)
		return nil, err
	}

	nodes := make([]neo4j.Node, 0)

	for result.Next() {
		record := result.Record()
		node, ok := record.GetByIndex(0).(neo4j.Node)
		if !ok {
			return nil, fmt.Errorf("failed to convert result to node")
		}
		nodes = append(nodes, node)
	}

	if err := result.Err(); err != nil {
		log.Printf("Failed to get nodes from result: %v", err)
		return nil, err
	}

	return nodes, nil
}
