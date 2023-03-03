# boostchickenlol
A clone of Metas bunnylol service written in go, configuration is via a YAML file, check the example in the repo.

Usage:
1. Deploy the docker container and expose the port you configured 
1. Create a custom search engine in your browser of choice 
```
Default: https://127.0.0.1:6969/%s
Traefik Proxy: https://lol.boostchicken.dev/%s
```
<img width="441" align="right" alt="Some people are curious what a boostchicken is.  Now you know." src="https://user-images.githubusercontent.com/427295/222669555-3b222ab6-ff78-4ce4-8735-d7cd5c104c09.png">

## Docker Build
[![Publish Docker image](https://github.com/boostchicken/lol/actions/workflows/docker-image.yml/badge.svg)](https://github.com/boostchicken/lol/actions/workflows/docker-image.yml)

### Docker Hub
[boostchicken/lol:latest](https://hub.docker.com/r/boostchicken/lol)

### Github Repo
[ghcr.io/boostchicken/lol:latest](https://github.com/boostchicken/lol/pkgs/container/lol)


Mount your config.yaml to /go/config.yaml as a volume
