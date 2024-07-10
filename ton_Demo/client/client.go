package client

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

var Client *ton.APIClient

func NewClient() {
	godotenv.Load()
	url := os.Getenv("testnet")
	client := liteclient.NewConnectionPool()
	err := client.AddConnectionsFromConfigUrl(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}
	Client = ton.NewAPIClient(client)
}
