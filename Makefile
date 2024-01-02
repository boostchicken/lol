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
	cd api && pnpm link .  && cd ../ui && pnpm link @boostchicken/lol-api &&  pnpm run build
.PHONY:ui
protoc: 
	protoc --go_out=src/model --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=src/model --proto_path=src/protoc lolconfig.proto

build: ui

	 cd ./src/cmd/lol/ && go build -ldflags="-s -w" -o ../../../bin/lol 
.PHONY:build
debugui:
	cd ui && pnpm dev
doc: build
	godoc
.PHONY:doc
