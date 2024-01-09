.PHONY: build
build:
	go build -o ./out/server -v ./cmd/server 

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

DEFAULT_GOAL := build
