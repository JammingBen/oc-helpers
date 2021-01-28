all: help

build:
	@go build -o build/oc-helpers ./cmd/oc-helpers.go

tidy:
	go mod tidy
	gofmt -s -w .

help:
	@echo "make reset-oc-data: reset all your oc data"
