.PHONY: build
build:
	go build -o ./out/server -v ./cmd/server 

.PHONY: build_nix
build_nix:
	GOOS=linux go build -o ./out/server_nix -v ./cmd/server 

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

DEFAULT_GOAL := build
