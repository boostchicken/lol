 # boostchicken - lol
A clone of Meta's bunnylol service written in go.

## Features
1. A UI and accompnaying REST endpoints to add or delete commands on the fly
2. Chrome Extension for querying your commands
3. History (last 250 queries and their result)
4. Config auto generation, only ned tyo change if you want to run on a different port than 8080
   
## Roadmap
1. Passkey Auth
2. Kube support, Multi Tenancy via k8s
3. Helm chart

   
### Usage
1. Deploy the docker container and expose the port you configured 
1. Create a custom search engine in your browser of choice 
```
Default: https://127.0.0.1:8080/lol?q=%s
Reverse Proxy: https://lol.boostchicken.dev/lol?q=%s
```

### UI
The UI is capable of managing config.yaml with no rehash or restarts. You can add and delete commnds on the fly
```
Default: https://127.0.0.1:8080/
Reverse Proxy: https://lol.boostchicken.dev/
```

![image](https://github.com/boostchicken/lol/assets/427295/950d2659-78be-463f-8a81-192baa65c2a6)

## Chrome Extension

<img width="315" alt="image" src="https://github.com/boostchicken/lol/assets/427295/4b5a23bf-d623-4cc1-8514-e01e687e25aa">

* Search your curent liveconfig for commands.

* In omnibox just type "bl " and the search term and the results / suggestions flow like above.
### Host config command
```bl setUrl http://<yourshost>/liveconfig```

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


### Config REST API 

*DELETE* /delete/:command

```Delete config entry by command```
* ```:command``` LoL Command

*PUT* /add/:command/:type?url=:url

```Add Command```

* ```:command``` LoL Command"
* ```:type``` one of [Alias,Redirect,RedirectVarArgs]
* ```:url``` printf format url

*Headers*
1. Content-Type 
   ```application/json```

Body
  * A config string that matches your header format (Yaml, JSON, TOML, etc)
```json
{
  "bind": "0.0.0.0:6969",
  "entries": [
    {
      "command": "g",
      "type": "Redirect",
      "value": "https://www.google.com/search?q=%s"
    },
    {
      "command": "github",
      "type": "RedirectVarArgs",
      "value": "https://www.github.com/%s/%s"
    },
    {
      "command": "pihole",
      "type": "Alias",
      "value": "http://pi.hole"
    }
  ]
}
 ```

*GET* /liveconfig

```Returns the current config in JSON```


* Content-Type: ```application/json```

*Body*
```json
{
  "bind": "0.0.0.0:6969",
  "entries": [
    {
      "command": "g",
      "type": "Redirect",
      "value": "https://www.google.com/search?q=%s"
    }
  ]
}
 ```

## Docker Build
[![Publish Docker image](https://github.com/boostchicken/lol/actions/workflows/docker-image.yml/badge.svg)](https://github.com/boostchicken/lol/actions/workflows/docker-image.yml)

### Docker Hub
[boostchicken/lol:latest](https://hub.docker.com/r/boostchicken/lol)

### Github Repo
[ghcr.io/boostchicken/lol:latest](https://github.com/boostchicken/lol/pkgs/container/lol)

### Development 
Use the provided devcontainer locally or in codespaces, there is a Makefile to build binaries, also the UI has a proxy config standard in Node

If you want to contribute but lack hardware for dev, I will add you to the repo and cover your codespaces cost.  Open an Issue if neeeded.

### Config Mounting example
```-v /tmp/config.yaml:/go/config.yaml```
* No longer needed unless you want to change the port, just boot the image and use the UI

### Easter Egg. 
* Find the REAL Boostchicken
* You can of course look at the source and cheat, but don't be lame
* First person to tell me the color of it gets a reward!
