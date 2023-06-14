package main

import (
	"log"

	"github.com/makaires77/iapf/backend/graph_pasteur/graphdb"
	"github.com/makaires77/iapf/backend/graph_pasteur/pkg"
)

func main() {
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

	// Inicializar o scraper
	scraper := pkg.NewScraper(neo4jClient)

	// Testar o scraping do website
	err = scraper.ScrapeWebsite()
	if err != nil {
		log.Fatalf("Failed to scrape website: %v", err)
	}

	// Testar a persistência dos dados no Neo4j
	err = scraper.PersistDataInNeo4j()
	if err != nil {
		log.Fatalf("Failed to persist data in Neo4j: %v", err)
	}

	log.Println("Scraping and data persistence completed successfully.")
}
