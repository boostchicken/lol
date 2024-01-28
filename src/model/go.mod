module model // import "github.com/boostchicken/lol/model"

go 1.21.6

require (
	github.com/infobloxopen/protoc-gen-gorm v1.1.2
	google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240125205218-1f4bbc51befe // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
)

replace (
	github.com/boostchicken/lol/clients/secrets => ../clients/secrets
	github.com/boostchicken/lol/config => ../config
)
