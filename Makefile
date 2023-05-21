.PHONY: generate
generate:
	buf generate

.PHONY: install
install:
	brew install bufbuild/buf/buf
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0
	go install github.com/evanw/esbuild/cmd/esbuild@v0.14.38
	cd frontend && pnpm i

.PHONY: serve
serve:
	cd backend && go run main.go