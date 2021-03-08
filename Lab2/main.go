package main

import (
	"context"
	"fib/fibonacci"
	"fib/kafka"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	if err := initGlobalTracer(nil); err != nil {
		panic(err)
	}
	/* Lab2-d: EXPORTING METRICS - TODO comment
	initMetricsProvider()
	*/
	http.HandleFunc("/fib", FibServer)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":28080", nil)
}

// FibServer handles HTTP requests for fibonacci calculation
func FibServer(w http.ResponseWriter, r *http.Request) {
	/* Lab2-a: CREATING INSTRUMENTATION LIBRARIES - TODO comment
	tracer := otel.Tracer("http")
	ctx := context.Background()

	var span trace.Span
	ctx, span = tracer.Start(ctx, "http-request")
	defer span.End()
	*/

	if n, err := getIntParam(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		///* Lab2-a: CREATING INSTRUMENTATION LIBRARIES - TODO un-comment
		result, numIterations := fibonacci.New().Calc(n)
		//*/

		/* Lab2-a: CREATING INSTRUMENTATION LIBRARIES - - TODO comment
		result, numIterations := fibonacci.Wrap(ctx, fibonacci.New()).Calc(n)
		*/
		reportMetric(n, numIterations)
		kafka.Send(result)
		w.Write([]byte(fmt.Sprintf("%d", result)))
	}
}

/*
DO NOT REMOVE ANY TEXT BELOW THIS LINE






















*/
func hide(v interface{}) {
	kafka.PreserveImport()
	fibonacci.PreserveImport()
	context.Background()
	var tracer trace.Tracer
	tracer = otel.Tracer("")
	hide(tracer)
	hide(fmt.Sprintf("%d", 3))
}
