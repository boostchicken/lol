package main

import (
	"github.com/boostchicken/lol/model"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE command = @command
	FilterWithCommand(command string) ([]gen.T, error)

	// SELECT * FROM @@table WHERE tenant = @tenant
	GetByTenant(tenant string) (*gen.T, error)
}

func main() {

	var dsn string = "postgresql://lol-dev:ZcCkF5rm2__SwksA7-Y4ww@boost-lol-764.j77.cockroachlabs.cloud:26257/dev?sslmode=verify-full"

	var Db, _ = gorm.Open(postgres.Open(dsn))
	Db.AutoMigrate(&model.LolEntry{})
	Db.AutoMigrate(&model.Config{})
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(Db) // reuse your gorm db

	g.ApplyBasic(model.Config{}, model.LolEntry{})
	g.ApplyInterface(func(Querier) {},model.Config{}, model.LolEntry{})

	// Generate the code
	g.Execute()
}
