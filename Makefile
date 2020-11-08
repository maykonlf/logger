SHELL := /bin/bash

SRC = $(shell find . -type f -name '*.go' -not -iname '*.pb.*' -not -path '**/mocks/*')

lint:
	@golangci-lint run ./...

test:
	@go test -cover ./...

bench:
	@go test -cpu 1,2,4,8 -benchmem -bench .

cover:
	@go test -cover -coverprofile=cover.out ./pkg/... > /dev/null
	@go tool cover -func cover.out

clean:
	@go clean

fmt:
	@gofmt -s -l -w $(SRC)

goimports:
	@goimports -w -local github.com/golangci/golangci-lint $(SRC)

format: fmt goimports

.PHONY: all test
