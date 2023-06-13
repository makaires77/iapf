package main

import (
	"fmt"
	"log"

	"github.com/makaires77/iapf/graphscraperpasteur/internal/graphdb"
	"github.com/makaires77/iapf/graphscraperpasteur/pkg"
)

func main() {
	// Configurações do Neo4j
	neo4jURI := "bolt://localhost:7687"
	neo4jUsername := "neo4j"
	neo4jPassword := "password"

	// Inicialização do cliente Neo4j
	neo4jClient, err := graphdb.NewNeo4j(neo4jURI, neo4jUsername, neo4jPassword)
	if err != nil {
		log.Fatalf("Failed to initialize Neo4j: %v", err)
	}

	defer neo4jClient.Close()

	// Executa o scraping do site e persiste os dados no Neo4j
	err = scrapeAndPersistData(neo4jClient)
	if err != nil {
		log.Fatalf("Failed to scrape and persist data: %v", err)
	}

	log.Println("Scraping and data persistence completed successfully!")
}

func scrapeAndPersistData(neo4jClient *graphdb.Neo4j) error {
	// Realiza o scraping do site
	scrapedData, err := pkg.ScrapeWebsite()
	if err != nil {
		return fmt.Errorf("failed to scrape website: %w", err)
	}

	// Itera sobre os dados raspados e persiste no Neo4j
	for _, data := range scrapedData {
		nodeProps := map[string]interface{}{
			"title": data.Title,
			"url":   data.URL,
		}

		_, err := neo4jClient.CreateNode("NodeLabel", nodeProps)
		if err != nil {
			return fmt.Errorf("failed to create node in Neo4j: %w", err)
		}
	}

	return nil
}
