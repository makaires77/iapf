package pkg

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/makaires77/iapf/backend/graph_pasteur/graphdb"
)

type Scraper struct {
	Neo4jClient *graphdb.Neo4j
}

func NewScraper(neo4jClient *graphdb.Neo4j) *Scraper {
	return &Scraper{
		Neo4jClient: neo4jClient,
	}
}

func (s *Scraper) ScrapeWebsite() error {
	// Realizar o scraping do site
	res, err := http.Get("https://research.pasteur.fr/en/")
	if err != nil {
		return fmt.Errorf("failed to make GET request to website: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Extrair o primeiro nível de nós
	headerInfos := doc.Find("#header-infos")
	postType := headerInfos.Find("div.post-type").Text()
	fmt.Println("Node:", postType)

	headerNav := headerInfos.Find("div.header-nav")
	items := headerNav.Find("div.item")
	items.Each(func(i int, item *goquery.Selection) {
		link := item.Find("a")
		count := link.Find("div.info div.count").Text()
		text := link.Find("div.info div.text").Text()
		fmt.Printf("Child Node %d: Count=%s, Text=%s\n", i+1, count, text)

		// Criar o nó filho e relacionamento no Neo4j
		childNodeID, err := s.Neo4jClient.CreateNode(text, map[string]interface{}{
			"count": count,
		})
		if err != nil {
			fmt.Printf("Failed to create child node: %v\n", err)
			return
		}

		// Criar relacionamento com o nó pai
		_, err = s.Neo4jClient.CreateRelationship(0, childNodeID, "HAS_CHILD", nil)
		if err != nil {
			fmt.Printf("Failed to create relationship: %v\n", err)
			return
		}

		// Recursivamente percorrer os nós filhos dos filhos
		s.ScrapeChildNode(link.AttrOr("href", ""), childNodeID)
	})

	// Processar os dados raspados
	// TODO: Criar nós e relacionamentos no Neo4j para os dados extraídos

	return nil
}

func (s *Scraper) ScrapeChildNode(url string, parentNodeID int64) {
	// Realizar o scraping do nó filho
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to make GET request to child node: %v\n", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("Unexpected status code for child node: %d\n", res.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Printf("Failed to parse HTML for child node: %v\n", err)
		return
	}

	// Extrair os nós do nó filho
	childNodes := doc.Find(".child-node")
	childNodes.Each(func(i int, childNode *goquery.Selection) {
		link := childNode.Find("a")
		count := link.Find("div.count").Text()
		text := link.Find("div.text").Text()
		fmt.Printf("Child Node %d: Count=%s, Text=%s\n", i+1, count, text)

		// Criar o nó filho e relacionamento no Neo4j
		childNodeID, err := s.Neo4jClient.CreateNode(text, map[string]interface{}{
			"count": count,
		})
		if err != nil {
			fmt.Printf("Failed to create child node: %v\n", err)
			return
		}

		// Criar relacionamento com o nó pai
		_, err = s.Neo4jClient.CreateRelationship(parentNodeID, childNodeID, "HAS_CHILD", nil)
		if err != nil {
			fmt.Printf("Failed to create relationship: %v\n", err)
			return
		}

		// Recursivamente percorrer os nós filhos dos filhos
		s.ScrapeChildNode(link.AttrOr("href", ""), childNodeID)
	})
}

func (s *Scraper) PersistDataInNeo4j() error {
	// TODO: Implementar a lógica para persistir os dados no Neo4j
	// Utilize s.Neo4jClient para interagir com o banco de dados Neo4j
	startNodeID := int64(1)
	endNodeID := int64(2)
	relType := "HAS"
	relProperties := map[string]interface{}{
		"since": "2022-01-01",
	}
	_, err := s.Neo4jClient.CreateRelationship(startNodeID, endNodeID, relType, relProperties)
	if err != nil {
		return fmt.Errorf("failed to create relationship in Neo4j: %w", err)
	}

	return nil
}
