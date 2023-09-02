package main // import "github.com/boostchicken/cmd/lol"

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/boostchicken/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

// Debug: /app/lol debug will run in debug mode
// Release: /app/lol will run in release mode
// reads the file : config.yaml right next to the exe
func main() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		newConf := config.Config{
			Bind: "0.0.0.0:8080",
			Entries: []config.LOLEntry{
				{
					Command: "g",
					Type:    "Redirect",
					Value:   "https://www.google.com/search?q=%s",
				},
			},
		}
		configFile = newConf.WriteConfig()
	}

	err3 := yaml.Unmarshal(configFile, &config.CurrentConfig)
	if err3 != nil {
		log.Fatal("unable to read config", err)
	}
	if len(config.CurrentConfig.Bind) == 0 {
		config.CurrentConfig.Bind = "0.0.0.0:8080"
	}
	config.CurrentConfig.CacheConfig()
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(static.Serve("/", static.LocalFile("./ui/build/", true)))
	r.Use(static.Serve("/api", static.LocalFile("./ui/build/", true)))
	r.GET("/rehash", InvokeRehash).GET("/config", RenderConfig).GET("/liveconfig", RenderConfigJSON).GET("/lol", Invoke).PUT("/add/:command/:type", AddCommand).DELETE("/delete/:command", DeleteCommand)
	r.GET("/history", RenderHistory)

	log.Println("Listening on", config.CurrentConfig.Bind)

	err4 := r.Run(config.CurrentConfig.Bind)
	if err4 != nil && err4 != http.ErrServerClosed {
		os.Exit(1)
	}
}

// DeleteCommand HTTP: GET /config
// Delete a current entry
// c gin.Context
func DeleteCommand(c *gin.Context) {
	for i, entry := range config.CurrentConfig.Entries {
		if entry.Command == c.Param("command") {

			config.CurrentConfig.Entries = append(config.CurrentConfig.Entries[:i], config.CurrentConfig.Entries[i+1:]...)

			config.CurrentConfig.CacheConfig()
			c.JSON(200, config.CurrentConfig)
			return
		}
	}
}

// RenderConfig HTTP: GET /config
// Renders current config based on Accept
// c gin.Context
func RenderConfig(c *gin.Context) {
	_ = c.BindHeader(config.CurrentConfig)
}

// RenderConfigJSON HTTP: GET /liveconfig
// Renders current config as JSON
// c gin.Context
func RenderConfigJSON(c *gin.Context) {
	c.JSON(200, config.CurrentConfig)
}

// RenderHistory HTTP: GET /history
// Renders current config as JSON
// c gin.Context
func RenderHistory(c *gin.Context) {
	c.JSON(200, maps.Values(config.HistoryCache.GetALL(false)))
}

// AddCommand HTTP: PUT /add/:command/:type?url=github.com
// c gin.Context
// Adds a new command and saves
func AddCommand(c *gin.Context) {
	typevar := c.Param("type")
	switch typevar {
	case "Redirect":
		break
	case "RedirectVarArgs":
		break
	case "Alias":
		break
	default:
		_ = c.AbortWithError(501, errors.New("Invalid type"))
	}
	config.CurrentConfig.Entries = append(config.CurrentConfig.Entries, config.LOLEntry{
		Command: strings.ToLower(strings.TrimSpace(c.Param("command"))),
		Type:    typevar,
		Value:   c.Query("url")})
	config.CurrentConfig.WriteConfig()
	config.CurrentConfig.CacheConfig()
	c.JSON(200, config.CurrentConfig)
}

// InvokeRehash HTTP: GET /rehash
// Renders current config as YAML
func InvokeRehash(c *gin.Context) {
	config.CurrentConfig.RehashConfig()
	c.YAML(200, gin.H{
		"message": "Rehashed",
	})
}

var t config.LOLAction = config.LOLAction{}

// Invoke HTTP: GET /lol?q=github boostchicken lol
// c gin.Context
// Query: q the actual command to run, space delimited
func Invoke(c *gin.Context) {
	command, ok := c.Params.Get("command")
	if !ok {
		q, qok := c.GetQuery("q")
		if !qok {
			c.YAML(501, gin.H{
				"message": "No command provided",
			})
			return
		}
		t.LOL(q, c)
		return
	}
	t.LOL(command, c)
}
