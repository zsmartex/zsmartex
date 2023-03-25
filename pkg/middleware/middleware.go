package middleware

import (
	"context"
	"net/http"
)

type contextKey int

var (
	requestContextKey contextKey = 0
)

func GetRequestFromContext(ctx context.Context) *http.Request {
	return ctx.Value(requestContextKey).(*http.Request)
}
