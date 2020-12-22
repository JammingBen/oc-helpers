all: help

reset-oc-data:
	@go run cmd/oc-helpers/reset-oc.go

tidy:
	go mod tidy
	gofmt -s -w .

help:
	@echo "make reset-oc-data: reset all your oc data"