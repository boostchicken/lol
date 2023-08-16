 # boostchicken - lol
A clone of Meta's bunnylol service written in go, configuration is via a YAML file, check the example in the repo.

### Usage
1. Deploy the docker container and expose the port you configured 
1. Create a custom search engine in your browser of choice 
```
Default: https://127.0.0.1:8080/lol?q=%s
Reverse Proxy: https://lol.boostchicken.dev/lol?q=%s
```

### UI
There is a UI capable of managing the config.yaml (Additions and Deletions) as well as a recent History of queries
```
Default: https://127.0.0.1:8080/
Reverse Proxy: https://lol.boostchicken.dev/
```

### Chrome Extension
As your configs grow its hard to remember everything this omnibox extension queries your live config and shows you the command and its url to sprintf to
 <img width="315" alt="image" src="https://github.com/boostchicken/lol/assets/427295/aae501ab-40ba-41c5-badd-a4b215b6db11">

### Config Types
* Alias - A straight redirect to a URL, no arguments used. (e.g.)
```yaml 
- command: gh
  type: Alias
  value: https://www.github.com
``` 

gh -> http://www.github.com
 
* Redirect - A redirect that accepts one argument to printf

```yaml 
- command: ghuser
  type: Redirect
  value: https://www.github.com/%s
``` 
ghuser boostchicken -> https://www.github.com/boostchicken

* RedirectVarArgs - A redirect that explodes all params on space and does allows multiple args into printf
```yaml 
- command: ghr
  type: RedirectVarArgs
  value: https://www.github.com/%s/%s
``` 

ghr boostchicken lol -> https://www.github.com/boostchicken/lol

### Config
*GET* /rehash 

```Reloads the configuration.  For use after editing config via filesystem```


*PUT* /config 

```Replaces the configuration file in memory, does not write it to disk.  Rehashes server after update.  This will not cause the server to rebind```


*Headers*
1. Content-Type 
   ```application/yaml```

Body
  * A config string that matches your header format (Yaml, JSON, TOML, etc)
```yaml
---
bind: "0.0.0.0:6969"
entries: 
  - command: g
    type: Redirect
    value: https://www.google.com/search?q=%s
  - command: github
    type: RedirectVarArgs
    value: https://www.github.com/%s/%s
  - command: pihole
    type: Alias   
    value: http://pi.hole
 ```

*Headers*

*GET* /config

```Returns the current config in YAML```


* Content-Type: ```application/yaml```

*Body*
```yaml
---
bind: "0.0.0.0:6969"
entries: 
  - command: g
    type: Redirect
    value: https://www.google.com/search?q=%s
 ```

## Docker Build
[![Publish Docker image](https://github.com/boostchicken/lol/actions/workflows/docker-image.yml/badge.svg)](https://github.com/boostchicken/lol/actions/workflows/docker-image.yml)

### Docker Hub
[boostchicken/lol:latest](https://hub.docker.com/r/boostchicken/lol)

### Github Repo
[ghcr.io/boostchicken/lol:latest](https://github.com/boostchicken/lol/pkgs/container/lol)


### Config Mounting example

```-v /tmp/config.yaml:/go/config.yaml```
