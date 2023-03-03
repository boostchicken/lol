package config

import (
	"log"
	"os"
	"reflect"

	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type LOLAction struct {
}

type LOLEntry struct {
	Command string
	Type    string
	Value   string
}

type Config struct {
	Bind    string
	Entries []LOLEntry
}

var CurrentConfig Config
var Cache map[string]LOLEntry
var ReflectionCache map[string]reflect.Method

func (c *Config) RehashConfig() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("unable to read config", err)
	}
	err = yaml.Unmarshal(configFile, &CurrentConfig)
	if err != nil {
		log.Fatal("unable to read config", err)
	}
	c.CacheConfig()
}

func (c *Config) CacheConfig() {
	Cache = make(map[string]LOLEntry)
	ReflectionCache = make(map[string]reflect.Method)

	for _, e := range c.Entries {
		Cache[e.Command] = e
		_, okm := ReflectionCache[e.Type]
		if !okm {

			method, okr := reflect.TypeOf(&LOLAction{}).MethodByName(e.Type)
			if !okr {
				log.Fatalf("Unable to find function %s", e.Type)
			}
			ReflectionCache[e.Type] = method
		}
	}
}

func (t *LOLAction) Redirect(c *gin.Context, url string, parts []string) {
	redir := fmt.Sprintf(url, strings.Join(parts[1:], " "))
	c.Redirect(http.StatusFound, redir)
}

func (t *LOLAction) Alias(c *gin.Context, url string, _ []string) {
	c.Redirect(http.StatusMovedPermanently, url)
}

func (t *LOLAction) RedirectVarArgs(c *gin.Context, url string, a ...any) {
	redir := fmt.Sprintf(url, a...)
	c.Redirect(http.StatusFound, redir)
}

func (t *LOLAction) LOL(command string, c *gin.Context) {
	explode := strings.Split(command, " ")
	entry, ok := Cache[explode[0]]
	parts := explode[1:]
	if !ok {
		if google, search := Cache["g"]; search {
			redir := fmt.Sprintf(google.Value, strings.Join(parts, " "))
			c.Redirect(http.StatusFound, redir)
		} else {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("no endpoint found"))
			return
		}
	}

	m, mok := ReflectionCache[entry.Type]
	if !mok {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("no endpoint found"))
		return
	}

	if strings.Contains(entry.Type, "VarArgs") {

		vars := explode[1:]
		var new = make([]interface{}, len(vars))
		for i, v := range vars {
			new[i] = interface{}(v) // or new[i] = v
		}
		m.Func.CallSlice([]reflect.Value{
			reflect.ValueOf(t),
			reflect.ValueOf(c),
			reflect.ValueOf(strings.TrimSpace(entry.Value)),
			reflect.ValueOf(new),
		})
	} else {
		m.Func.Call([]reflect.Value{
			reflect.ValueOf(t),
			reflect.ValueOf(c),
			reflect.ValueOf(strings.TrimSpace(entry.Value)),
			reflect.ValueOf(parts),
		})
	}

}
