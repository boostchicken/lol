package config

import (
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_RedirectVarArgs(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		action LOLAction
	}{{
		"RedirectVarArgs",
		Config{
			Bind: "0.0.0.0:6969",
			Entries: []LOLEntry{
				{
					Command: "github",
					Type:    "RedirectVarArgs",
					Value:   "https://www.github.com/%s/%s",
				}},
		},
		LOLAction{},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Config = tt.config
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			c.CacheConfig()
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
