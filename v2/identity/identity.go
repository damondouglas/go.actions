package identity

import "net/http"

// Identity generates http.HandlerFunc for OAuth2 requests.
type Identity interface {
	AuthorizationHandler(r *http.Request) (http.HandlerFunc, error)
	ExchangeHandler(r *http.Request) (http.HandlerFunc, error)
}
