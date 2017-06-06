package gobmx

import (
	"context"
	"log"
	"net/http"
)

type key int

const (
	keyBmxLog key = 1
)

func createLoggingContext(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, keyBmxLog, "JSON")
}

// LoggingHandler wrap your handler with the very handy logging handler
func LoggingHandler(h http.Handler, l *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := createLoggingContext(req.Context(), req)
		h.ServeHTTP(w, req.WithContext(ctx))
		l.Printf("%s %s\n", req.Method, req.URL)
	})
}

// ContextLogger a Context based logger that lets you log simple things like errors that will get printed into the weblog at the right time
type ContextLogger interface{}
