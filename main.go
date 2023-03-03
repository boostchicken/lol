package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/boostchicken/lol/config"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
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
		bytes, err := yaml.Marshal(newConf)
		if err != nil {
			log.Fatal("unable to write default config")
		}
		_ = os.WriteFile("config.yaml", bytes, fs.ModePerm)
		configFile = bytes

	}

	err = yaml.Unmarshal(configFile, &config.CurrentConfig)
	if err != nil {
		log.Fatal("unable to read config", err)
	}
	if config.CurrentConfig.Bind == "" {
		config.CurrentConfig.Bind = "0.0.0.0:8080"
	}
	config.CurrentConfig.CacheConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/rehash", InvokeRehash).GET("/:command", Invoke)
	log.Println("Listening on", config.CurrentConfig.Bind)

	err = r.Run(config.CurrentConfig.Bind)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("unable to start server", err)
	}
}

func InvokeRehash(c *gin.Context) {
	config.CurrentConfig.RehashConfig()
	c.YAML(200, gin.H{
		"message": "Rehashed",
	})
}

var t config.LOLAction = config.LOLAction{}

func Invoke(c *gin.Context) {
	command := c.Param("command")
	if c.Query("q") != "" {
		command = c.Query("q")
	}
	t.LOL(command, c)
}
