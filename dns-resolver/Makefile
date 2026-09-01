.PHONY: run test build

BINARY = dns-resolver

run:
	go run ./cmd $(DOMAIN)

test:
	go test ./...

build:
	go build -o $(BINARY) ./cmd

help:
	@echo "make run DOMAIN=google.com"
	@echo "make test"
	@echo "make build"
