package test

import (
	"log"
	"testing"

	"github.com/makaires77/iapf/backend/graph_pasteur/graphdb"
	"github.com/stretchr/testify/assert"
)

func TestNeo4j(t *testing.T) {
	// Configuração do Neo4j
	uri := "bolt://localhost:7687"
	username := "neo4j"
	password := "password"

	// Inicializar o cliente do Neo4j
	neo4jClient, err := graphdb.NewNeo4j(uri, username, password)
	if err != nil {
		log.Fatalf("Failed to create Neo4j client: %v", err)
	}
	defer neo4jClient.Close()

	// Testar a criação de um nó
	nodeProps := map[string]interface{}{
		"name":  "Node 1",
		"value": 123,
	}
	node, err := neo4jClient.CreateNode("TestNode", nodeProps)
	assert.NoError(t, err)
	assert.NotNil(t, node)

	// Testar a criação de um relacionamento
	relProps := map[string]interface{}{
		"type":  "REL_TYPE",
		"value": "Relationship 1",
	}
	relationship, err := neo4jClient.CreateRelationship(node, node, "RELATED_TO", relProps)
	assert.NoError(t, err)
	assert.NotNil(t, relationship)

	// Testar a busca de nós
	matchProps := map[string]interface{}{
		"name": "Node 1",
	}
	nodes, err := neo4jClient.MatchNodes("TestNode", matchProps)
	assert.NoError(t, err)
	assert.NotNil(t, nodes)
	assert.NotEmpty(t, nodes)
}
