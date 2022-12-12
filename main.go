package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
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

type T struct{}

func main() {
	cache = make(map[string]LOLEntry)

	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		newConf := Config{
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
		ioutil.WriteFile("config.yaml", bytes, fs.ModePerm)
		configFile, _ = ioutil.ReadFile("config.yaml")

	}
	err = yaml.Unmarshal(configFile, &currentConfig)
	if err != nil {
		log.Fatal("unable to read config", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/{command}", InvokeLOL)
	http.ListenAndServe(currentConfig.Bind, r)
}

func InvokeLOL(w http.ResponseWriter, r *http.Request) {
	command := mux.Vars(r)["command"]
	parts := strings.Split(command, " ")
	entry, ok := cache[parts[0]]
	if !ok {
		for _, e := range currentConfig.Entries {
			cache[e.Command] = e
		}
		val, ok2 := cache[parts[0]]
		if !ok2 {
			w.WriteHeader(404)
			return
		}
		entry = val
	}
	var t T
	m, okm := reflect.TypeOf(&t).MethodByName(entry.Type)
	if !okm {
		w.WriteHeader(404)
		return
	}
	var in = make([]reflect.Value, 5)
	in[0] = reflect.ValueOf(&t)
	in[1] = reflect.ValueOf(w)
	in[2] = reflect.ValueOf(r)
	in[3] = reflect.ValueOf(strings.TrimSpace(entry.Value))
	in[4] = reflect.ValueOf(parts)

	m.Func.Call(in)
}

func (t *T) Redirect(w http.ResponseWriter, r *http.Request, url string, parts []string) {
	redir := fmt.Sprintf(url, strings.Join(parts[1:], " "))
	w.Header().Add("Location", redir)
	w.WriteHeader(302)
}

func (t *T) Alias(w http.ResponseWriter, r *http.Request, url string, parts []string) {
	w.Header().Add("Location", url)
	w.WriteHeader(302)
}
func (t *T) RedirectVarArgs(w http.ResponseWriter, r *http.Request, url string, parts ...string) {
	redir := fmt.Sprintf(url, parts)
	w.Header().Add("Location", redir)
	w.WriteHeader(302)
}
