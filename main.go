package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
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

type actions struct{}

var t actions

func main() {

	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		newConf := Config{
			Bind: ":8080",
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
		os.WriteFile("config.yaml", bytes, fs.ModePerm)
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

	r := mux.NewRouter()
	r.HandleFunc("/rehash", InvokeRehash).Queries()
	r.HandleFunc("/{command}", InvokeLOL)
	log.Println("Listening on", currentConfig.Bind)

	err = http.ListenAndServe(currentConfig.Bind, r)
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

func InvokeRehash(w http.ResponseWriter, r *http.Request) {
	currentConfig.RehashConfig()
}

func InvokeLOL(w http.ResponseWriter, r *http.Request) {
	command := mux.Vars(r)["command"]
	if r.FormValue("q") != "" {
		command = r.FormValue("q")
	}
	parts := strings.Split(command, " ")
	entry, ok := cache[parts[0]]
	if !ok {
        if google, search := cache["g"]; search {
            redir := fmt.Sprintf(google.Value, strings.Join(parts, " "))
            http.Redirect(w, r, redir, http.StatusFound)
        } else {
            http.NotFound(w, r)
            return
        }
	}

	m, mok := reflectionCache[entry.Type]
	if !mok {
		http.NotFound(w, r)
		return
	}

	m.Func.Call([]reflect.Value{
		reflect.ValueOf(&t),
		reflect.ValueOf(w),
		reflect.ValueOf(r),
		reflect.ValueOf(strings.TrimSpace(entry.Value)),
		reflect.ValueOf(parts),
	})
}

func (t *actions) Redirect(w http.ResponseWriter, r *http.Request, url string, parts []string) {
	redir := fmt.Sprintf(url, strings.Join(parts[1:], " "))
	http.Redirect(w, r, redir, http.StatusFound)
}

func (t *actions) Alias(w http.ResponseWriter, r *http.Request, url string, parts []string) {
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func (t *actions) RedirectVarArgs(w http.ResponseWriter, r *http.Request, url string, parts ...string) {
	redir := fmt.Sprintf(url, parts)
	http.Redirect(w, r, redir, http.StatusFound)
}
