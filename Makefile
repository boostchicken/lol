.DEFAULT_GOAL := build
export GO11MODULES=on
fmt:
	go fmt ./src/cmd/lol/*.go
	go fmt ./src/internal/config/*.go
.PHONY:fmt

lint: fmt
	golint ./src/cmd/lol/*.go
	golint ./src/internal/config/*.go
.PHONY:lint

vet: fmt
	cd ./src/cmd/lol && go vet main.go
.PHONY:vet

build: vet
	cd ./src/cmd/lol  && go build -ldflags="-s -w" -o ../../../bin/lol
.PHONY:build

doc: build
	godoc
.PHONY:doc
