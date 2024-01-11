package config

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boostchicken/lol/model"
	"github.com/gin-gonic/gin"
)

func Test_RedirectVarArgs(t *testing.T) {
	tests := []struct {
		name   string
		config model.Config
		action LOLAction
	}{{
		name: "RedirectVarargs",
		config: model.Config{
			Tenant: "boost",
			Bind:   "0.0.0.0:6969",
			Entries: []*model.LolEntry{
				{
					Command: "github",
					Type:    model.CommandType_Redirect,
					Url:     "https://www.github.com/%s/%s",
				}},
		},
		action: LOLAction{},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c model.Config = tt.config
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			url := "http://localhost:6969/lol?q=github boostchicken lol"
			request, _ := http.NewRequest(http.MethodGet, url, nil)
			request.RequestURI = url
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = request
			CurrentConfig = c
			CacheConfig()
			tt.action.LOL("github boostchicken lol", ctx)
			log.Println(w.Result())
		})
	}
}

func Test_Sprint(t *testing.T) {
	vars := []string{"command", "go", "boostchicken"}[1:]
	var new = make([]interface{}, len(vars))
	for i, v := range vars {
		new[i] = interface{}(v) // or new[i] = v
	}
	println(fmt.Sprintf("go%s/%s", new...))
}
