module github.com/boostchicken/lol/query

go 1.21.6

require (
	github.com/boostchicken/lol/clients/secrets v0.0.0-20211013185905-8b9b8b8b8b8b
	github.com/boostchicken/lol/model v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.5.4
	gorm.io/gen v0.3.25
	gorm.io/gorm v1.25.5
	gorm.io/plugin/dbresolver v1.5.0
)

require (
	github.com/aws/aws-sdk-go-v2 v1.24.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.26.4 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.16.15 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.26.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.7 // indirect
	github.com/aws/smithy-go v1.19.0 // indirect
	github.com/boostchicken/lol/config v0.0.0-20230303115543-49611d96b059 // indirect
	github.com/bytedance/sonic v1.8.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/infobloxopen/protoc-gen-gorm v1.1.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.2 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.10 // indirect
	golang.org/x/arch v0.2.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	google.golang.org/genproto v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240116215550-a9fa1716bcac // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/datatypes v1.2.0 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	gorm.io/hints v1.1.2 // indirect
)

replace (
	github.com/boostchicken/lol/clients/secrets => ../clients/secrets
	github.com/boostchicken/lol/model => ../model
)
