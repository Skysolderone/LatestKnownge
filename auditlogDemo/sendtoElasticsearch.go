package main

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

func sendToEsticsearch() {
	cfg := elasticsearch.Config{Addresses: []string{"http://localhost:9200"}}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open("audit.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry AuditLogEntry
		if err := json.Unmarshal([]byte(scanner.Text()), &entry); err != nil {
			log.Fatal(err)
		}
		body, err := json.Marshal(entry)
		if err != nil {
			log.Fatal(err)
		}
		req := esapi.IndexRequest{
			Index:        "audit.log",
			DocumentType: "_doc",
			DocumentID:   "",
			Boyd:         strings.NewReader(string(body)),
			Refresh:      "true",
		}
		_, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatal(err)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
