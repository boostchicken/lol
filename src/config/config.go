package config //import "github.com/boostchicken/lol/config"

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/boostchicken/lol/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ops uint64
var wg sync.WaitGroup
var Db *gorm.DB
var err error

type LOLAction struct {
}

// CurrentConfig the current config loaded
var CurrentConfig model.Config

var cache map[string]*model.LolEntry = make(map[string]*model.LolEntry)         // A Map that caches LOLEntry BY Command
var reflectionCache map[string]reflect.Method = make(map[string]reflect.Method) // Caches the Method associated with the Type

// RehashConfig Reload the config but do not rebind

// CacheConfig Generate ReflectionCache and Command Cache
func CacheConfig() {
	for _, e := range CurrentConfig.Entries {
		cache[e.GetCommand()] = e
		_, okm := reflectionCache[e.GetType().String()]
		if !okm {

			method, okr := reflect.TypeOf(&LOLAction{}).MethodByName(e.GetType().String())
			if !okr {
				log.Fatalf("Unable to find function %s", e.Type)
			}
			reflectionCache[e.GetType().String()] = method
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
	//model.HistoryCache.Set(ops, model.History{Command: c.Query("q"), Result: res})
}

// AddCommandHistory add command execution to history cache
func (t *LOLAction) AddCommandHistory(result string, c *gin.Context) {
	wg.Add(1)
	defer wg.Done()
	ops++
	//history := model.History{Command: c.Query("q"), Result: result}

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
func (t *LOLAction) RedirectVarargs(c *gin.Context, url string, a ...any) {
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
			redir := fmt.Sprintf(google.GetUrl(), strings.Join(explode, " "))
			c.Redirect(http.StatusFound, redir)
			t.AddCommandHistory(redir, c)
			return
		}
		_ = c.AbortWithError(http.StatusNotFound, fmt.Errorf("Unable to find cache entry for %s", explode[0]))
	}

	m, mok := reflectionCache[entry.GetType().String()]
	if !mok {
		_ = c.AbortWithError(http.StatusNotFound, fmt.Errorf("unale to find reflection cache entry for %s", entry.Type))
		return
	}

	if strings.Contains(entry.GetType().String(), "Varargs") {
		vars := explode[1:]
		var new = make([]interface{}, len(vars))
		for i, v := range vars {
			new[i] = interface{}(v) // or new[i] = v
		}
		m.Func.CallSlice([]reflect.Value{
			reflect.ValueOf(t),
			reflect.ValueOf(c),
			reflect.ValueOf(strings.TrimSpace(entry.GetUrl())),
			reflect.ValueOf(new),
		})
	} else {
		m.Func.Call([]reflect.Value{
			reflect.ValueOf(t),
			reflect.ValueOf(c),
			reflect.ValueOf(strings.TrimSpace(entry.GetUrl())),
			reflect.ValueOf(explode[0:]),
		})
	}
}
