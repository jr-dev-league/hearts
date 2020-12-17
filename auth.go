package main

import (
	"context"
	"net/http"
)

type authCtxKey string

var authCtx = authCtxKey("auth")

// authorizeEndpoint is a middleware that figures out who the user is. For v1 it just checks the
// header and believes whatever it's told. In v2 it will have to do some actual authorizing.
func authorizeEndpoint(next http.HandlerFunc) (authHandler http.HandlerFunc) {
	authHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		header := req.Header
		auth := header.Get("authorization")

		if auth == "" {
			writeResponse(w, req, resError{"Forbidden"}, 403)

			return
		}

		ctx := context.WithValue(req.Context(), authCtx, user{auth})
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})

	return
}
