.PHONY: build test lint run

build:
	go build -o yml-generator ./cmd/main.go

test:
	go test -v ./...

lint:
	go vet ./...

run: build
	./yml-generator