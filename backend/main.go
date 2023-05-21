// Copyright 2023 Andy Challis. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"log"
	"net/http"

	usersv1 "github.com/ghandic/grpc-web-go-react-example/backend/gen/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/users"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func main() {
	gs := grpc.NewServer()
	usersv1.RegisterUserServiceServer(gs, &users.UserService{})
	wrappedServer := grpcweb.WrapServer(gs)

	http.Handle("/api/", http.StripPrefix("/api/", wrappedServer))

	log.Println("Serving on http://0.0.0.0:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
