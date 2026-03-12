// Package middleware contains middlewares to be added to the requests in the onion style.
package middleware

import "net/http"

// Middleware defines the exact signature you requested.
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares to a http.HandlerFunc.
func Chain(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	// Loop backwards to ensure middlewares are executed in the order they are passed.
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
