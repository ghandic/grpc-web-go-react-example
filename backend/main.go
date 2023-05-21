// Copyright 2023 Andy Challis. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1/usersv1connect"
	"github.com/ghandic/grpc-web-go-react-example/backend/users"
)

func main() {

	mux := http.NewServeMux()

	userService := &users.UserService{}
	path, handler := usersv1connect.NewUserServiceHandler(userService)
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
