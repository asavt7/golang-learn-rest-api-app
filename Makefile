.PHONY: build

build:
	go build -v -o ./bin/main ./cmd/main.go


.DEFAULT_GOAL := build