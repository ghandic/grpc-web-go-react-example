// Copyright 2023 Andy Challis. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1/usersv1connect"
	"github.com/ghandic/grpc-web-go-react-example/backend/users"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	zapadapter "github.com/jackc/pgx-zap"
)

func newCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}

func getPGPool() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	logger, err := zap.NewDevelopmentConfig().Build()

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   zapadapter.NewLogger(logger),
		LogLevel: tracelog.LogLevelTrace,
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return pool
}

func main() {

	mux := http.NewServeMux()

	pool := getPGPool()
	defer pool.Close()

	userService := &users.UserService{
		Pool: pool,
	}
	path, handler := usersv1connect.NewUserServiceHandler(userService)
	mux.Handle(path, handler)

	err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(newCORS().Handler(mux), &http2.Server{}),
	)
	fmt.Fprintf(os.Stderr, "Unable to start server: %v\n", err)
	os.Exit(1)

}
