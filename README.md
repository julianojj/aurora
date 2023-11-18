<div align="center">
  <img src="https://github.com/julianojj/aurora/blob/develop/aurora.jpg" alt="aurora logo" />
</div>

# aurora

[![Go Report Card](https://goreportcard.com/badge/github.com/julianojj/aurora)](https://goreportcard.com/report/github.com/julianojj/aurora)

## Introduction
Aurora is a back end project made in Go that was based on Adobe XD. In it, good development practices were used, applying concepts of clean architecture, design patterns and tdd.

## Technologies
```
aws-sdk-go v1.46.3
gin v1.9.1
google/uuid v1.3.1
godotenv v1.5.1
minio-go/v7 v7.0.63
testify v1.8.4
zap v1.26.0
```

## Folder Structure
```
├── cmd
└── internal
    ├── config
    ├── core
    │   ├── domain
    │   ├── exceptions
    │   └── usecases
    └── infra
        ├── adapters
        ├── api
        │   ├── controllers
        │   ├── middlewares
        │   └── routes
        └── repository
```

cmd: have main file where is a aggregate root

internal/core: have business rules and application rules, have a domain entities, repository interface, exceptions and usecases
  
internal/infra: it's a layers of external access. Have adapters patters, repositories, api and database connection


## As run a project?
Before run project, should config env files, after configuration, your need run a docker-compose with `docker-compose up -d` to started a minio server

`go test ./...` run tests

`go run ./cmd/main.go` run api
