package config //import "github.com/boostchicken/lol/config"

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

// LOLAction is the main action struct
// Mainly needed for reflection
type LOLAction struct {
}

// Configuration entry
type LOLEntry struct {
	Command string // the first arguement of q delimited by spaces
	Type    string // Any of the three types. RedirectVarArgs, Alias, Redirect
	Value   string // the url to redirect to ex: https://www.google.com/search?q=%s
}

// Config is the main config struct
type Config struct {
	Bind    string     // HTTP Bind address
	Entries []LOLEntry // List of entries
}

var CurrentConfig Config                      // the current config
var Cache map[string]LOLEntry                 // A Map that caches LOLEntry BY Command
var ReflectionCache map[string]reflect.Method // Caches the Method associated with the Type

// Reload the config but do not rebind
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

// Generate ReflectionCache and Command Cache
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

// Simple printf http:://google.com/search?q=%s
// c gin context
// url the url as a string
// parts command split by spaces
func (t *LOLAction) Redirect(c *gin.Context, url string, parts []string) {
	redir := fmt.Sprintf(url, strings.Join(parts[1:], " "))
	c.Redirect(http.StatusFound, redir)
}

// A static redirect
// c gin context
// url the url as a string
// _ a pointless variable i let making reflectiosn easier
func (t *LOLAction) Alias(c *gin.Context, url string, _ []string) {
	c.Redirect(http.StatusMovedPermanently, url)
}

// A VarArgs redirect http:://github.com/%s/%s
// c gin context
// url the url as a string
// a varargs for printf
func (t *LOLAction) RedirectVarArgs(c *gin.Context, url string, a ...any) {
	redir := fmt.Sprintf(url, a...)
	c.Redirect(http.StatusFound, redir)
}

// Find the command and then execute the function associated with the Type
// command the command to execute
// c gin context
func (t *LOLAction) LOL(command string, c *gin.Context) {
	explode := strings.Split(command, " ")
	entry, ok := Cache[explode[0]]
	if !ok {
		if google, search := Cache["g"]; search {
			redir := fmt.Sprintf(google.Value, strings.Join(explode, " "))
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
			reflect.ValueOf(explode[0:]),
		})
	}

}
