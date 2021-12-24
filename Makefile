.PHONY: build
build:
	   go build -v ./cmd/k_app

.DEFAULT_GOAL := build