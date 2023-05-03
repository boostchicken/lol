 # boostchickenlol
A clone of Metas bunnylol service written in go, configuration is via a YAML file, check the example in the repo.

### Usage
1. Deploy the docker container and expose the port you configured 
1. Create a custom search engine in your browser of choice 
```
Default: https://127.0.0.1:6969/?q=%s
Traefik Proxy: https://lol.boostchicken.dev/lol?q=%s
```

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
<img width="441" align="right" alt="Some people are curious what a boostchicken is.  Now you know." src="https://user-images.githubusercontent.com/427295/222669555-3b222ab6-ff78-4ce4-8735-d7cd5c104c09.png">

## Docker Build
[![Publish Docker image](https://github.com/boostchicken/lol/actions/workflows/docker-image.yml/badge.svg)](https://github.com/boostchicken/lol/actions/workflows/docker-image.yml)

### Docker Hub
[boostchicken/lol:latest](https://hub.docker.com/r/boostchicken/lol)

### Github Repo
[ghcr.io/boostchicken/lol:latest](https://github.com/boostchicken/lol/pkgs/container/lol)


### Config Mounting example

```-v /tmp/config.yaml:/go/config.yaml```
