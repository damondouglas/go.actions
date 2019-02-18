package helper

import (
	"context"
	"net/http"
)

// ContextFactory generates context.Context from http.Request.
type ContextFactory interface {
	NewContext(r *http.Request) context.Context
}

// ClientFactory generates http.Client.
type ClientFactory interface {
	NewClient(context.Context) *http.Client
}
