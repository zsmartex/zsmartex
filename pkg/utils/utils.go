package utils

import (
	"context"
	"net/http"
)

type ContextKey int

const (
	RequestContextKey ContextKey = iota
)

func GetRequestFromContext(ctx context.Context) *http.Request {
	return ctx.Value(RequestContextKey).(*http.Request)
}

func SetRequestToHeader(ctx context.Context, request *http.Request) context.Context {
	return context.WithValue(ctx, RequestContextKey, request)
}
