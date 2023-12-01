package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	clientsOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientsOptions)
	if err != nil {
		log.Fatal("connect:", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("ping:", err)
	}
	fmt.Println("Contect to mongodb")
}
