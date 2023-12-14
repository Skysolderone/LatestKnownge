package main

import (
	"fmt"
	"log"
	"net/http"
)

//
//docker build -t my-demo-app .
// docker run -p 8080:8080 my-demo-app

//deployment
//kubectl apply -f deployment.yaml
//kubectl apply -f service.yaml

// kubectl get deployments
// kubectl get services
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is Shantanu!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
