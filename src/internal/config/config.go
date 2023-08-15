package config //import "github.com/boostchicken/lol/config"

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/bluele/gcache"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

var ops uint64
var wg sync.WaitGroup

// HistoryCache LRU cache for command history
var HistoryCache = gcache.New(250).LRU().Build()

// History entry struct
type History struct {
	Command   string
	Result    string
	ipAddress string
}

// LOLAction is the main action struct
// Mainly needed for reflection
type LOLAction struct {
}

// LOLEntry Configuration entry
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

// CurrentConfig the current config loaded
var CurrentConfig Config

var cache map[string]LOLEntry                 // A Map that caches LOLEntry BY Command
var reflectionCache map[string]reflect.Method // Caches the Method associated with the Type

// RehashConfig Reload the config but do not rebind
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

// WriteConfig write config to config.yaml
func (c *Config) WriteConfig() []byte {
	bytes, err2 := yaml.Marshal(&CurrentConfig)
	if err2 != nil {
		log.Fatal("unable to write default config")
	}
	_ = os.WriteFile("config.yaml", bytes, fs.ModePerm)
	return bytes
}

// CacheConfig Generate ReflectionCache and Command Cache
func (c *Config) CacheConfig() {
	cache = make(map[string]LOLEntry)
	reflectionCache = make(map[string]reflect.Method)

	for _, e := range c.Entries {
		cache[e.Command] = e
		_, okm := reflectionCache[e.Type]
		if !okm {

			method, okr := reflect.TypeOf(&LOLAction{}).MethodByName(e.Type)
			if !okr {
				log.Fatalf("Unable to find function %s", e.Type)
			}
			reflectionCache[e.Type] = method
		}
	}
}

// Redirect  Simple printf http:://google.com/search?q=%s
// c gin context
// url the url as a string
// parts command split by spaces
func (t *LOLAction) Redirect(c *gin.Context, url string, parts []string) {
	res := fmt.Sprintf(url, strings.Join(parts[1:], " "))
	c.Redirect(http.StatusFound, res)
	t.AddCommandHistory(res, c)
}

// AddCommandHistory add command execution to history cache
func (t *LOLAction) AddCommandHistory(result string, c *gin.Context) {
	wg.Add(1)
	ops++
	HistoryCache.Set(ops, History{Command: c.Query("q"), Result: result, ipAddress: c.ClientIP()})
	wg.Done()
}

// Alias A static redirect
// c gin context
// url the url as a string
// _ a pointless variable i let making reflections easier
func (t *LOLAction) Alias(c *gin.Context, url string, _ []string) {
	res := url
	c.Redirect(http.StatusMovedPermanently, res)
	t.AddCommandHistory(res, c)
}

// RedirectVarArgs A VarArgs redirect http:://github.com/%s/%s
// c gin context
// url the url as a string
// a varargs for printf
func (t *LOLAction) RedirectVarArgs(c *gin.Context, url string, a ...any) {
	res := fmt.Sprintf(url, a...)
	c.Redirect(http.StatusFound, res)
	t.AddCommandHistory(res, c)
}

// LOL Find the command and then execute the function associated with the Type
// command the command to execute
// c gin context
func (t *LOLAction) LOL(command string, c *gin.Context) {
	explode := strings.Split(command, " ")
	entry, ok := cache[explode[0]]
	if !ok {
		if google, search := cache["g"]; search {
			redir := fmt.Sprintf(google.Value, strings.Join(explode, " "))
			c.Redirect(http.StatusFound, redir)
			t.AddCommandHistory(redir, c)
			return
		}
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("no endpoint found"))
	}

	m, mok := reflectionCache[entry.Type]
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
