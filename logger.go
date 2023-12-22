package main

import (
	"context"
	"os"

	"github.com/hirosassa/zerodriver"
)

var logger = zerodriver.NewProductionLogger()

// info インフォログの出力
func info(ctx context.Context, msg string) {
	span := SpanFromContext(ctx)
	logger.Info().
		TraceContext(span.SpanContext().TraceID().String(), span.SpanContext().SpanID().String(), true, os.Getenv("PROJECT_ID")).
		Msg(msg)
}
