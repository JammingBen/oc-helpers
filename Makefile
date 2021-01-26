all: help

oc-helpers:
	@go run ./cmd/oc-helpers.go

tidy:
	go mod tidy
	gofmt -s -w .

help:
	@echo "make reset-oc-data: reset all your oc data"
