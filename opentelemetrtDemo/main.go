package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

const portNum = ":8080"

var tracer = otel.Tracer("info-service")

type InfoResponse struct {
	Version     string `json:"version"`
	ServiceName string `json:"service-name"`
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := InfoResponse{Version: "0.1.0", ServiceName: "otlp-sample"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	log.Println("Starting http server.")

	mux := http.NewServeMux()
	ctx := context.Background()
	consoleTraceExporter, err := newTraceExporter()
	if err != nil {
		log.Println("Failed get console exporter.")
	}

	tracerProvider := newTraceProvider(consoleTraceExporter)

	defer tracerProvider.Shutdown(ctx)
	otel.SetTracerProvider(tracerProvider)

	mux.HandleFunc("/info", info)
	srv := &http.Server{
		Addr:    portNum,
		Handler: mux,
	}

	log.Println("Started on port", portNum)
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Fail start http server.")
	}
}

func newTraceExporter() (trace.SpanExporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

func newTraceProvider(traceExporter trace.SpanExporter) *trace.TracerProvider {
	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			trace.WithBatchTimeout(time.Second)),
	)
	return traceProvider
}
