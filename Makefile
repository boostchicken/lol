.DEFAULT_GOAL := build
export GO111MODULES=on
fmt: protoc
	gofmt -w ./src
.PHONY:fmt
lint: fmt
	 golangci-lint run ./src/query/... ./src/model/... ./src/cmd/... ./src/api/... ./src/ui/... ./src/... ./src/... ./src/...
	 .PHONY:lint
ui: protoc
	cd ui && pnpm link ../api &&  pnpm run build
.PHONY:ui
protoc:
	protoc -Isrc/protos/google/api -Isrc/protos --go_out=./src/model --go_opt=paths=source_relative --go-grpc_out=./src/model --go-grpc_opt=paths=source_relative gorm.proto lolconfig.proto
PHONY:protoc
build: ui
	 cd ./src/cmd/lol && go build -ldflags="-s -w" -o ../../../bin/lol 
.PHONY:  build
debugui:
	cd ui && pnpm dev
.PHONY:debugui
