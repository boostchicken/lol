package query

import (
	"fmt"

	"github.com/boostchicken/lol/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func create() {
	var dsn string = "postgresql://lol-dev:ZcCkF5rm2__SwksA7-Y4ww@boost-lol-764.j77.cockroachlabs.cloud:26257/dev?sslmode=verify-full"
	var Db, _ = gorm.Open(postgres.Open(dsn))
	conf := &model.Config{Tenant: "dorman", Bind: "0.0.0.0:6969"}
	c := Config.Where(Config.Bind.Neq("1"))
	db, err := c.Find()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(db[len(db)-1].GetBind())
	newConfig(Db).Create(conf)
	newLolEntry(Db).Create(&model.LolEntry{Config: conf, Command: "g", Url: "https://www.google.com/q=%s", Type: model.CommandType_Redirect})
}
