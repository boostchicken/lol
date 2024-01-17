.DEFAULT_GOAL := build
export GO11MODULES=on
fmt: protoc
	gofmt -w ./src
.PHONY:fmt
lint: fmt
	 golangci-lint run ./src/query/... ./src/model/... ./src/cmd/... ./src/api/... ./src/ui/... ./src/... ./src/... ./src/...
	 
.PHONY:lint
ui: protoc
	cd ./src/ui && pnpm link ../api &&  pnpm run build
.PHONY:ui
protoc:
	protoc --go_out=./src/model --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=./src/model --proto_path=./protos openapiv2.proto annotations.proto gorm.proto  lolconfig.proto
PHONY:protoc
build: ui
	 cd ./src/cmd/ && go build -ldflags="-s -w" -o ../../../bin/lol 

debugui:
	cd ui && pnpm dev
.PHONY:debugui
