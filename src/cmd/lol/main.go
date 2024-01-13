package main // import "github.com/boostchicken/cmd/lol"

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/boostchicken/internal/config"
	"github.com/boostchicken/lol/model"
	"github.com/boostchicken/lol/query"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var Q *query.Query

func init() {
	Q = query.Use(lolconf.GetDb())
}

// Debug: /app/lol debug will run in debug mode
// Release: /app/lol will run in release mode
// reads the file : config.yaml right next to the exe
func main() {
	c := Q.Config
	models, err := c.WithContext(context.Background()).Where(c.Tenant.Eq("dorman")).Find()
	if err != nil && len(models) == 0 {
		newConf := model.Config{
			Tenant: "dorman",
			Bind:   "0.0.0.0:8080",
			Entries: []*model.LolEntry{
				{
					Command: "g",
					Type:    model.CommandType_Redirect,
					Url:     "https://www.google.com/search?q=%s",
				},
			},
		}
		err2 := c.Create(&newConf)
		if err2 != nil {
			log.Fatal("unable to create config", err2)
		}

	if len(config.CurrentConfig.Bind) == 0 {
		*config.CurrentConfig.Bind = "0.0.0.0:8080"
	}
	config.CacheConfig()
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	fs := static.LocalFile("./ui/out/", true)
	r := gin.Default()
	r.Use(cors.Default(), static.Serve("/api", fs), static.Serve("/", fs))

	r.GET("/config", RenderConfig).GET("/liveconfig", RenderConfigJSON)
	r.GET("/lol", Invoke).GET("/history", RenderHistory)
	r.PUT("/add/:command/:type", AddCommand).DELETE("/delete/:command", DeleteCommand)



	err4 := r.Run(config.CurrentConfig.Bind)
	if err4 != nil && err4 != http.ErrServerClosed {
		log.Fatal("unable to start server", err4)
	}
}

// DeleteCommand HTTP: GET /config
// Delete a current entry
// c gin.Context
func DeleteCommand(c *gin.Context) {
	l := Q.LolEntry
	command, err := l.Where(l.ConfigId.Eq(config.CurrentConfig.Id), l.Command.Eq(c.Param("command"))).First()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	l.Delete(command)

	for i, entry := range config.CurrentConfig.Entries {
		if entry.Command == c.Param("command") {

			config.CurrentConfig.Entries = append(config.CurrentConfig.Entries[:i], config.CurrentConfig.Entries[i+1:]...)

			config.CacheConfig()
		
		}
	}
	c.JSON(200, &config.CurrentConfig)
	
}

// RenderConfig HTTP: GET /config
// Renders current config based on Accept
// c gin.Context
func RenderConfig(c *gin.Context) {
	c.ProtoBuf(200, &config.CurrentConfig)
}

// RenderConfigJSON HTTP: GET /liveconfig
// Renders current config as JSON
// c gin.Context
func RenderConfigJSON(c *gin.Context) {
	c.JSON(200, &config.CurrentConfig)
}

// RenderHistory HTTP: GET /history
// Renders current config as JSON
// c gin.Context
func RenderHistory(c *gin.Context) {
	c.JSON(200, model.HistoryList{})
}

// AddCommand HTTP: PUT /add/:command/:type?url=github.com
// c gin.Context
// Adds a new command and saves
func AddCommand(c *gin.Context) {
	typevar := c.Param("type")
	var enumType model.CommandType
	switch typevar {
	case "Redirect":
		enumType = model.CommandType_Redirect
	case "RedirectVarags":
		enumType = model.CommandType_RedirectVarargs
	case "Alias":
		enumType = model.CommandType_Alias
	default:
		_ = c.AbortWithError(501, errors.New("invalid type"))
	}
	config.CurrentConfig.Entries = append(config.CurrentConfig.Entries, &model.LolEntry{
		Config:  &config.CurrentConfig,
		Command: strings.ToLower(strings.TrimSpace(c.Param("command"))),
		Type:    enumType,
		Url:     c.Query("url")})

	Q.LolEntry.Create(config.CurrentConfig.Entries...)
	config.CacheConfig()
	c.JSON(200, &config.CurrentConfig)
}

var t config.LOLAction = config.LOLAction{}

// Invoke HTTP: GET /lol?q=github boostchicken lol
// c gin.Context
// Query: q the actual command to run, space delimited
func Invoke(c *gin.Context) {
	q, qok := c.GetQuery("q")
	if !qok {
		c.JSON(501, gin.H{
			"message": "No command provided",
		})
		return
	}
	t.LOL(q, c)
}
