package model // import "github.com/boostchicken/lol/model"

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/boostchicken/lol/clients/secrets"
	"github.com/boostchicken/lol/config"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE command = @command
	FilterWithCommand(command string) ([]gen.T, error)
}

func main() {

	dsn, err := secrets.GetDSN()
	if err != nil {
		panic(err)
	}
	Db, err := gorm.Open(postgres.Open(aws.ToString(dsn)), &gorm.Config{})

	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithQueryInterface,
	})

	g.UseDB(Db)
	g.GenerateAllTable()
	g.ApplyBasic(config.Config{}, config.LolEntry{})

	g.ApplyInterface(func(Querier) {}, config.LolEntry{})

	g.Execute()
}

func file_gorm_proto_init() {
	return
}
