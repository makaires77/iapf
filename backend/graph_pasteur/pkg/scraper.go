package pkg

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"golang.org/x/net/html"
)

func ScrapeWebsiteAndPersistToNeo4j(driver neo4j.Driver) {
	// Criação do coletor
	c := colly.NewCollector()

	// Callback para páginas HTML
	c.OnHTML("html", func(e *colly.HTMLElement) {
		processNode(driver, e.DOM.Nodes[0], "")
	})

	// Callback para erros de requisição
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Algo deu errado:", err)
	})

	// Iniciar o scraping
	c.Visit("https://www.research.pasteur.fr/en/")
}

func processNode(driver neo4j.Driver, n *html.Node, parent string) {
	switch n.Type {
	case html.ElementNode:
		nodeType := n.Data
		props := make(map[string]interface{})
		for _, a := range n.Attr {
			props[a.Key] = a.Val
		}
		createNodeInNeo4j(driver, nodeType, props)
		parent = nodeType
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			createNodeInNeo4j(driver, "text", map[string]interface{}{"value": text})
			createRelationshipInNeo4j(driver, parent, "text")
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processNode(driver, c, parent)
	}
}

func createNodeInNeo4j(driver neo4j.Driver, nodeType string, props map[string]interface{}) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(
			fmt.Sprintf("CREATE (n:%s $props) RETURN n", nodeType),
			map[string]interface{}{"props": props},
		)
	})

	if err != nil {
		log.Printf("Failed to create node: %v", err)
		return
	}

	if result.Next() {
		fmt.Printf("Created node: %v\n", result.Record().GetByIndex(0))
	}
}

func createRelationshipInNeo4j(driver neo4j.Driver, from string, to string) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(
			fmt.Sprintf("MATCH (a:%s), (b:%s) WHERE id(a) < id(b) CREATE (a)-[r:RELATED_TO]->(b)", from, to),
			nil,
		)
	})

	if err != nil {
		log.Printf("Failed to create relationship: %v", err)
	}
}
