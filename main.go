package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type LOLEntry struct {
	Command string
	Type    string
	Value   string
}

type Config struct {
	Bind    string
	Entries []LOLEntry
}

var currentConfig Config
var cache map[string]LOLEntry
var reflectionCache map[string]reflect.Method

type actions struct {
}

var t actions

func main() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		newConf := Config{
			Bind: "0.0.0.0:8080",
			Entries: []LOLEntry{
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

	err = yaml.Unmarshal(configFile, &currentConfig)
	if err != nil {
		log.Fatal("unable to read config", err)
	}
	if currentConfig.Bind == "" {
		currentConfig.Bind = "0.0.0.0:8080"
	}
	currentConfig.CacheConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/rehash", InvokeRehash).GET("/:command", InvokeLOL)
	log.Println("Listening on", currentConfig.Bind)

	err = r.Run(currentConfig.Bind)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("unable to start server", err)
	}
}

func (c *Config) RehashConfig() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("unable to read config", err)
	}
	err = yaml.Unmarshal(configFile, &currentConfig)
	if err != nil {
		log.Fatal("unable to read config", err)
	}
	c.CacheConfig()
}

func (c *Config) CacheConfig() {
	cache = make(map[string]LOLEntry)
	reflectionCache = make(map[string]reflect.Method)

	for _, e := range c.Entries {
		cache[e.Command] = e
		_, okm := reflectionCache[e.Type]
		if !okm {

			method, okr := reflect.TypeOf(&t).MethodByName(e.Type)
			if !okr {
				log.Fatalf("Unable to find function %s", e.Type)
			}
			reflectionCache[e.Type] = method
		}
	}
}

func InvokeRehash(c *gin.Context) {
	currentConfig.RehashConfig()
	c.YAML(200, gin.H{
		"message": "Rehashed",
	})
}

func InvokeLOL(c *gin.Context) {
	command := c.Param("command")
	if c.Query("q") != "" {
		command = c.Query("q")
	}
	parts := strings.Split(command, " ")
	entry, ok := cache[parts[0]]
	if !ok {
		if google, search := cache["g"]; search {
			redir := fmt.Sprintf(google.Value, strings.Join(parts, " "))
			c.Redirect(http.StatusFound, redir)
		} else {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("no endpoint found"))
			return
		}
	}

	m, mok := reflectionCache[entry.Type]
	if !mok {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("no endpoint found"))
		return
	}

	m.Func.Call([]reflect.Value{
		reflect.ValueOf(&t),
		reflect.ValueOf(c),
		reflect.ValueOf(strings.TrimSpace(entry.Value)),
		reflect.ValueOf(parts),
	})
}

func (t *actions) Redirect(c *gin.Context, url string, parts []string) {
	redir := fmt.Sprintf(url, strings.Join(parts[1:], " "))
	c.Redirect(http.StatusFound, redir)
}

func (t *actions) Alias(c *gin.Context, url string, _ []string) {
	c.Redirect(http.StatusMovedPermanently, url)
}

func (t *actions) RedirectVarArgs(c *gin.Context, url string, parts ...string) {
	redir := fmt.Sprintf(url, parts)
	c.Redirect(http.StatusFound, redir)
}
