.PHONY: generate
generate:
	buf generate --debug
	cd backend/db/connectors && sqlc generate

.PHONY: install
install:
	go install github.com/cosmtrek/air@latest
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@v1.7.0
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	pnpm i -g @bufbuild/protoc-gen-connect-query @bufbuild/protoc-gen-es

	cd frontend && pnpm i

.PHONY: serve/backend
serve/backend:
	cd backend && air

.PHONY: serve/frontend
serve/frontend:
	cd frontend && pnpm run dev

.PHONY: serve/db
serve/db:
	docker-compose up

.PHONY: test/backend
test/backend:
	go test ./...