package query

import (
	"fmt"

	"github.com/boostchicken/lol/clients/secrets"
	"github.com/boostchicken/lol/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func create() {

	dsn, err := secrets.GetDSN()
	if err != nil {
		panic(err)
	}
	db, err2 := gorm.Open(postgres.Open(*dsn))
	if err2 != nil {
		panic(err2)
	}
	conf := &model.Config{Tenant: "dorman", Bind: "0.0.0.0:6969"}
	c := Config.Where(Config.Bind.Neq("1"))
	dbq, err3 := c.Find()
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	fmt.Println(dbq[len(dbq)-1].GetBind())
	newLolEntry(db).Create(&model.LolEntry{Config: conf, Command: "gx", Url: "https://www.google.com/q=%s", Type: model.CommandType_Redirect})
}
