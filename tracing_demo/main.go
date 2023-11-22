package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// docker run \
// --rm \
// -p 6831:6831/udp \
// -p 6832:6832/udp \
// -p 16686:16686 \
// jaegertracing/all-in-one:1.7 \
// --log-level=debug
func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	return tracer, closer
}
func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	helloTo := os.Args[1]
	//init tracing
	// tracer := opentracing.GlobalTracer()
	tracer, closer := initJaeger("hello-demo")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("say-hello")
	defer span.Finish()
	span.SetTag("hello-to", helloTo)
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	helloStr := formatString(ctx, helloTo)
	// span.LogFields(
	// 	log.String("event", "string-format"),
	// 	log.String("value", helloStr),
	// )

	printHello(ctx, helloStr)
	// fmt.Println(helloStr)
	// span.LogKV("event", "println")

}
func formatString(ctx context.Context, s string) string {
	//rootspan := span.Tracer().StartSpan("formatString", opentracing.ChildOf(span.Context()))
	rootspan, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer rootspan.Finish()
	helloStr := fmt.Sprintf("Hello, %s!", s)
	rootspan.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	return helloStr
}
func printHello(ctx context.Context, s string) {
	//rootspan := span.Tracer().StartSpan("printhello")
	rootspan, _ := opentracing.StartSpanFromContext(ctx, "printhello")
	defer rootspan.Finish()
	println(s)
	rootspan.LogKV("event", "println")
}
