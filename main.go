package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

	"github.com/boostchicken/lol/config"
)

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
		bytes, err2 := yaml.Marshal(newConf)
		if err2 != nil {
			log.Fatal("unable to write default config")
		}
		_ = os.WriteFile("config.yaml", bytes, fs.ModePerm)
		configFile = bytes
	}

	err3 := yaml.Unmarshal(configFile, &config.CurrentConfig)
	if err3 != nil {
		log.Fatal("unable to read config", err)
	}
	if len(config.CurrentConfig.Bind) == 0 {
		config.CurrentConfig.Bind = "0.0.0.0:8080"
	}
	config.CurrentConfig.CacheConfig()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/rehash", InvokeRehash).GET("/config", RenderConfigYAML).GET("/", Invoke).GET("/:command").PUT("/config", updateConfig)
	log.Println("Listening on", config.CurrentConfig.Bind)

	err4 := r.Run(config.CurrentConfig.Bind)
	if err4 != nil && err4 != http.ErrServerClosed {
		log.Fatal("unable to start server", err)
	}
}

func RenderConfigYAML(c *gin.Context) {
	c.YAML(200, config.CurrentConfig)
}

func updateConfig(c *gin.Context) {
	c.BindHeader(config.CurrentConfig)
	config.CurrentConfig.RehashConfig()
	c.YAML(200, gin.H{"message": "Updated"})
}

func InvokeRehash(c *gin.Context) {
	config.CurrentConfig.RehashConfig()
	c.YAML(200, gin.H{
		"message": "Rehashed",
	})
}

var t config.LOLAction = config.LOLAction{}

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
