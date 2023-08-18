# go-packcal
Micro-DDD (Hex-Arc) RestAPI and Web UI for Package Calculator of Shipping

[![Build](https://github.com/husamettinarabaci/go-packcal/actions/workflows/build.yml/badge.svg)](https://github.com/husamettinarabaci/go-packcal/actions/workflows/build.yml)
[![Test](https://github.com/husamettinarabaci/go-packcal/actions/workflows/test.yml/badge.svg)](https://github.com/husamettinarabaci/go-packcal/actions/workflows/test.yml)
[![Test](https://github.com/husamettinarabaci/go-packcal/actions/workflows/sast.yml/badge.svg)](https://github.com/husamettinarabaci/go-packcal/actions/workflows/sast.yml)

<a href="https://kaos.sh/g/go-badge"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev" /></a>
<a href="https://kaos.sh/w/go-badge/ci"><img src="https://kaos.sh/w/go-badge/ci.svg" alt="GitHub Actions CI Status" /></a>
<a href="https://github.com/husamettinarabaci/go-packcal/stargazers"><img src="https://img.shields.io/github/stars/husamettinarabaci/go-packcal" alt="Stars Badge"/></a>
<a href="https://github.com/husamettinarabaci/go-packcal/network/members"><img src="https://img.shields.io/github/forks/husamettinarabaci/go-packcal" alt="Forks Badge"/></a>
<a href="https://github.com/husamettinarabaci/go-packcal/pulls"><img src="https://img.shields.io/github/issues-pr/husamettinarabaci/go-packcal" alt="Pull Requests Badge"/></a>
<a href="https://github.com/husamettinarabaci/go-packcal/issues"><img src="https://img.shields.io/github/issues/husamettinarabaci/go-packcal" alt="Issues Badge"/></a>
<a href="https://github.com/husamettinarabaci/go-packcal/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/husamettinarabaci/go-packcal?color=2b9348"></a>
<a href="https://github.com/husamettinarabaci/go-packcal/blob/master/LICENSE"><img src="https://img.shields.io/github/license/husamettinarabaci/go-packcal?color=2b9348" alt="License Badge"/></a>

## Stack
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)

## Dependencies
[![Go-Gin](https://img.shields.io/badge/GoLib-Gin-green.svg)](https://github.com/gin-gonic/gin/)
[![Go-Ioc](https://img.shields.io/badge/GoLib-Ioc-green.svg)](https://github.com/golobby/container/v3/)
[![Go-Json](https://img.shields.io/badge/GoLib-Json-green.svg)](https://github.com/goccy/go-json/)
[![Go-Yaml](https://img.shields.io/badge/GoLib-Yaml-green.svg)](https://gopkg.in/yaml.v3/)
[![Go-Uuid](https://img.shields.io/badge/GoLib-Uuid-green.svg)](https://github.com/google/uuid/)
[![Testify](https://img.shields.io/badge/GoLib-Testify-green.svg)](https://github.com/stretchr/testify/)

## Getting Started
<b>PackCal</b> provides optimal configuration of package size and product count for shipping with RestAPI or Web UI. You can run it locally or as a container. 

Live Demo: [https://packcal.husamettinarabaci.com](https://packcal.husamettinarabaci.com)

It has been developed with <b>Domain Driven Design (Hex-Arc)</b> architecture and allows you to be included in the domain and perform external operations in all microservice infrastructures without additional development processes. 

Do you want to learn more information about <b>Domain Driven Design</b> and <b>Hex-Arc</b>? 
 - [DDD](https://en.wikipedia.org/wiki/Domain-driven_design)
 - [Hex-Arc](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))

Also, you can see these pictures about <b>Domain Driven Design (Hex-Arc)</b>. 
 - [Hex-Arc-1](https://github.com/husamettinarabaci/go-packcal/tree/main/doc/Hex-Arc-1.jpg)
 - [Hex-Arc-2](https://github.com/husamettinarabaci/go-packcal/tree/main/doc/Hex-Arc-2.jpg)


## Installation
```bash
git clone https://github.com/husamettinarabaci/go-packcal.git
cd go-packcal
go mod download
```

## Test
```bash
go test -v ./...
```

## Usage
You can access the Web Site with the below URL.
```bash
http://localhost:16000
```

or

You can use with default package sizes. 
Default package sizes are 250, 500, 1000, 2000, 5000, 10000, 20000, 50000.
```bash
POST http://localhost:16080/api/calc HTTP/1.1
content-type: application/json

{
    "item": 250
}
```

or

You can use with custom package sizes. 
```bash
POST http://localhost:16080/api/calc HTTP/1.1
content-type: application/json

{
    "item": 250
    "pack_sizes": [40, 50, 60, 70]
}
```

## Local Run
```bash
export LOCAL=true && go run cmd/main.go
```

## Docker Build & Run
```bash
docker build -t {DOCKER_USERNAME}/{YOUR_REPO} -f script/Dockerfile .
docker tag {DOCKER_USERNAME}/{YOUR_REPO} {DOCKER_USERNAME}/{YOUR_REPO}:latest
docker push {DOCKER_USERNAME}/{YOUR_REPO}:latest
docker run -p 16080:16080 -p 16000:16000 {DOCKER_USERNAME}/{YOUR_REPO}:latest
```

## Kubernetes Deploy
```bash
kubectl apply -f script/k8s.yml
```

## Github Actions
Fork the project and create the below secrets in your repo.

```bash
DOCKERHUB_USERNAME

DOCKERHUB_TOKEN
```

Create "release" branch and create a pull request to "release" branch and merge it. Github Actions will build and push docker image to your dockerhub repo.

## Project Structure - Domain Driven Design (Hex-Arc)
```bash
.
├── cmd
│   └── main.go
├── config
│   ├── log_local.yml
│   ├── log.yml
│   ├── rest_local.yml
│   ├── rest.yml
│   ├── web_local.yml
│   └── web.yml
├── core
│   ├── application
│   │   ├── infrastructure
│   │   │   └── port
│   │   │       ├── calc.go
│   │   │       └── log.go
│   │   ├── presentation
│   │   │   ├── adapter
│   │   │   │   ├── calc.go
│   │   │   │   └── query.go
│   │   │   └── port
│   │   │       ├── command.go
│   │   │       └── query.go
│   │   └── service
│   │       └── service.go
│   └── domain
│       ├── model
│       │   ├── entity
│       │   │   ├── calcrequest.go
│       │   │   └── calcresponse.go
│       │   ├── interface
│       │   │   └── loggable.go
│       │   └── object
│       │       ├── calc.go
│       │       ├── error.go
│       │       └── response.go
│       └── service
│           ├── service.go
│           └── service_test.go
├── doc
│   ├── Hex-Arc-1.jpg
│   └── Hex-Arc-2.jpg
├── go.mod
├── go.sum
├── LICENSE
├── pkg
│   ├── infrastructure
│   │   ├── adapter
│   │   │   ├── calc.go
│   │   │   └── log.go
│   │   └── mapper
│   │       └── calc.go
│   └── presentation
│       ├── controller
│       │   ├── rest
│       │   │   └── restapi.go
│       │   └── web
│       │       ├── views
│       │       │   └── index.html
│       │       └── web.go
│       └── dto
│           ├── calcrequest.go
│           └── calcresponse.go
├── README.md
├── script
│   ├── Dockerfile
│   └── k8s.yml
└── tool
    ├── config
    │   ├── log.go
    │   ├── rest.go
    │   └── web.go
    ├── json
    │   └── json.go
    └── slice
        └── slice.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Tasks
- [ ] Redesign the project WebUI.

## Contact

Hüsamettin ARABACI - info@husamettinarabaci.com

Project Link: [https://github.com/husamettinarabaci/go-packcal](https://github.com/husamettinarabaci/go-packcal)

