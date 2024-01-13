package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/boostchicken/lol/model"
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

	dsnInput := sm.database.{
		SecretId: aws.String("boost-lol-dev")}

	dsn := sm.DescribeSecretInput(dsnInput).

	Db = &gorm.DB{postgres.Open(aws.String(dsn.SecretString))}

	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(Db)
	g.GenerateAllTable()
	g.ApplyBasic(model.Config{}, model.LolEntry{})
	g.ApplyInterface(func(Querier) {}, model.LolEntry{})

	g.Execute()
}
