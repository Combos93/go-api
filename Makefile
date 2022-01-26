.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: cover
cover:
	go test -coverprofile=coverage.out -count 1 ./...
	go tool cover -html=coverage.out

.DEFAULT_GOAL := build
