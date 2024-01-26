package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//docker run \
// -p 9090:9090 \
// -v /path/to/prometheus.yml:/etc/prometheus/prometheus.yml \
// prom/prometheus

//docker run -p 9090:9090 prom/prometheus


//grafana
//docker run -d --name=grafana -p 3000:3000 grafana/grafana-enterprise

func main() {

	// Create a new registry.
	reg := prometheus.NewRegistry()

	// Add Go module build info.
	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	fmt.Println("Hello world from new Go Collector!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
