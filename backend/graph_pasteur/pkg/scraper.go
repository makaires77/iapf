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
	res, err := http.Get("http://www.research.pasteur.fr/en/")
	if err != nil {
		return fmt.Errorf("failed to make GET request to website: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	_, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Processar os dados raspados
	// TODO: Extrair as informações desejadas e criar nós e relacionamentos no Neo4j

	return nil
}

func (s *Scraper) PersistDataInNeo4j() error {
	// TODO: Implementar a lógica para persistir os dados no Neo4j
	// Utilize s.Neo4jClient para interagir com o banco de dados Neo4j

	return nil
}
