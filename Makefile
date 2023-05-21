.PHONY: generate
generate:
	buf generate

.PHONY: install
install:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	pnpm i - g @bufbuild/protoc-gen-connect-query @bufbuild/protoc-gen-es
	cd frontend && pnpm i

.PHONY: serve/backend
serve/backend:
	air

.PHONY: serve/frontend
serve/frontend:
	cd frontend && pnpm run dev

.PHONY: serve/db
serve/db:
	docker-compose up