package graphdb

import (
	"fmt"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Neo4j struct {
	Driver neo4j.Driver
}

func NewNeo4j(uri, username, password string) (*Neo4j, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, fmt.Errorf("error creating neo4j driver: %w", err)
	}

	return &Neo4j{
		Driver: driver,
	}, nil
}

func (n *Neo4j) Close() error {
	return n.Driver.Close()
}

func (n *Neo4j) ExecuteQueryWithRetry(query string, parameters map[string]interface{}) (*neo4j.Result, error) {
	var result *neo4j.Result
	var err error

	for i := 0; i < 3; i++ {
		result, err = n.ExecuteQuery(query, parameters)
		if err == nil {
			return result, nil
		}

		// Exponential backoff
		time.Sleep(time.Second * time.Duration((i+1)*(i+1)))
	}

	return result, err
}

func (n *Neo4j) ExecuteQuery(query string, parameters map[string]interface{}) (*neo4j.Result, error) {
	session, _ := n.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(query, parameters)
	})

	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return nil, err
	}

	return result.(*neo4j.Result), nil
}

func (n *Neo4j) CreateNode(label string, properties map[string]interface{}) (*neo4j.Result, error) {
	query := fmt.Sprintf("CREATE (n:%s $props) RETURN n", label)
	return n.ExecuteQueryWithRetry(query, map[string]interface{}{
		"props": properties,
	})
}

func (n *Neo4j) CreateRelationship(fromNodeID, toNodeID int64, relType string, properties map[string]interface{}) (*neo4j.Result, error) {
	query := fmt.Sprintf("MATCH (a), (b) WHERE id(a) = $fromID AND id(b) = $toID CREATE (a)-[r:%s $props]->(b) RETURN r", relType)
	return n.ExecuteQueryWithRetry(query, map[string]interface{}{
		"fromID": fromID,
		"toID":   toNodeID,
		"props":  properties,
	})
}

func (n *Neo4j) MatchNodes(label string, properties map[string]interface{}) (*neo4j.Result, error) {
	query := fmt.Sprintf("MATCH (n:%s) WHERE properties(n) = $props RETURN n", label)
	return n.ExecuteQueryWithRetry(query, map[string]interface{}{
		"props": properties,
	})
}
