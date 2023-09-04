.DEFAULT_GOAL := build
export GO11MODULES=on
fmt:
	go fmt ./src/cmd/lol/*.go
	go fmt ./src/internal/config/*.go
.PHONY:fmt

lint: fmt
	golangci-lint run src/cmd/lol/*.go
	golangci-lint  run src/internal/config/*.go
.PHONY:lint

vet: fmt
	cd ./src/cmd/lol && go vet main.go
.PHONY:vet
ui: vet
	 cd ui && npx next build
PHONY: ui
build: ui
	 cd ./src/cmd/lol/ && go build -ldflags="-s -w" -o ../../../bin/lol 
.PHONY:build

debugui:
	cd ui && npx next dev
		
doc: build
	godoc
.PHONY:doc