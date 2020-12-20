package main

import (
	"context"
	"net/http"
	"strconv"
)

// This is a typed string, it's being used to avoid collisions in ctx.Value
type authCtxKey string

// This is the value that will be used to get auth information from ctx.
var authCtx = authCtxKey("user")

// This is the user object that will be stored in ctx
type user struct {
	ID uint8
}

// authorizeEndpoint is a middleware that figures out who the user is. For v1 it just checks the
// header and believes whatever it's told. In v2 it will have to do some actual authorizing.
//
// A middleware takes a handlerFunc (next) and returns a new handlerFunc that does some logic before
// executing next. In this case, the logic it's adding is a check to make sure there is an
// authorization header. If it finds it, it adds it to the Request's context so it can be used by
// handler functions. If it does not find it it returns a 403.
func authorizeEndpoint(next http.HandlerFunc) (authHandler http.HandlerFunc) {

	// authHandler is the handler that will be returned. By wrapping next in this handler, we get
	// added functionality.
	authHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// Get the Authoriaztion header
		header := req.Header
		auth := header.Get("Authorization")

		// If there is no header...
		if auth == "" {
			writeResponse(w, req, resError{"Forbidden"}, 403)

			return
		}

		userID, err := strconv.ParseInt(auth, 10, 8)
		if err != nil {
			writeResponse(w, req, resError{"Forbidden"}, 403)

			return
		}

		// create a new context with the Request context as its parent. Add the Authorization header
		// value to it inside a user object.
		ctx := context.WithValue(req.Context(), authCtx, user{uint8(userID)})
		req = req.WithContext(ctx) // Replaces Request context with this new, enhanced context.

		// Run the given function
		next.ServeHTTP(w, req)
	})

	return
}

func getUser(req *http.Request) user {
	return req.Context().Value(authCtx).(user)
}
