package test

import (
	"log"
	"testing"

	es "github.com/elastic/go-elasticsearch/v8"
)

var (
	client *es.Client
)

// func init() {
// 	var err error
// 	client, err = es.NewClient(es.Config{
// 		Addresses: []string{"http://localhost:9200"},
// 		Username:  "username",
// 		Password:  "password",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
func TestEs(t *testing.T) {
	t.Log(client.Info())
}
