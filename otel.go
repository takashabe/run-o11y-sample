package main

import (
	"context"
	"log"
	"net/http"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

// InitTracer TraceProviderの初期化
func InitTracer(ctx context.Context) (shutdown func(context.Context) error) {
	projectID := os.Getenv("PROJECT_ID")
	otel.SetTextMapPropagator(propagation.TraceContext{})

	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		log.Fatalf("texporter.New: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("run-o11y-sample"),
		)),
	)
	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

// SpanFromHeader HTTPヘッダからトレース情報を取得
func SpanFromHeader(ctx context.Context, header http.Header) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(header))
}

// StartSpan Spanの記録の開始
func StartSpan(ctx context.Context, name string) context.Context {
	tr := otel.GetTracerProvider().Tracer("run-o11y-sample")
	cctx, _ := tr.Start(ctx, name)
	ctx = cctx
	return ctx
}

// SpanFromContext ContextからSpanを取得
func SpanFromContext(ctx context.Context) trace.Span {
	return trace.SpanFromContext(ctx)
}

// EndSpan Spanの記録の終了
func EndSpan(ctx context.Context) {
	trace.SpanFromContext(ctx).End()
}
