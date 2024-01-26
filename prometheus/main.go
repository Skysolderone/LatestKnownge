package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//docker run  -p 9090:9090  -v G:\LatestKnownge\prometheus\prometheus.yml:/etc/prometheus/prometheus.yml   prom/prometheus

var pingCounter = prometheus.NewTimer(
	prometheus.ObserverFunc{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func main() {
	// http.HandleFunc("/ping", ping)

	// http.ListenAndServe(":8090", nil)
	prometheus.MustRegister(pingCounter)
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8090", nil)
}
