package main

import (
	"context"
	"net/http"
)

func main() {
	ctx := context.Background()
	fn := InitTracer(ctx)
	defer func() {
		fn(ctx)
	}()

	http.HandleFunc("/", traceHandle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// traceHandle ヘッダからトレース情報を取得して、contextを伝播させる
func traceHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = SpanFromHeader(ctx, r.Header)
	defer func() { EndSpan(ctx) }()

	run(ctx)
	info(ctx, "log/traceHandle")
}

// run トレースとログを出力
func run(ctx context.Context) {
	ctx = StartSpan(ctx, "trace/run")
	defer func() { EndSpan(ctx) }()

	info(ctx, "log/run")
	run2(ctx)
}

// run2 トレースとログを出力その2
func run2(ctx context.Context) {
	ctx = StartSpan(ctx, "trace/run2")
	defer func() { EndSpan(ctx) }()

	info(ctx, "log/run2")
}
