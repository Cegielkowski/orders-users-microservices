package elastic

import (
	"fmt"
	"github.com/olivere/elastic"
)

// Get ElasticSearch Client.
func GetESClient() (*elastic.Client, error) {
	client, err :=  elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err
}