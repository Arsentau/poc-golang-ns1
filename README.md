# NS1 Golang Poc

## Overview

A Poc to interact with the NS1 API to retrieve and process Networks and Zones information.
Two approaches were taken:

- Consume the API with a custom service
- Consume the API using the official [SDK](https://github.com/ns1/ns1-go)

## How to run

### Install dependencies

`go mod tidy`

### Create .env file

Create a `.env` file in the root folder of the project using [.env.example](/.env.example) as reference.

### Run Executables

With the .env with all variables set, run positioned in the root folder one of the following commands:

- **Custom Client** run positioned in the root folder `go run cmd/custom_ns1_client_api/custom_ns1_client_api.go`
- **Official Client (SDK)** `go run cmd/official_ns1_client_api/official_ns1_client_api.go`

## Endpoints

### Custom Client

- `/zones`
- `/networks`

### Official Client (SDK)

- `/zones`

## Project Structure

```bash
├── cmd
│   ├── custom_ns1_client_api
│   │   └── custom_ns1_client_api.go
│   └── official_ns1_client_api
│       └── official_ns1_client_api.go
├── go.mod
├── go.sum
├── internal
│   └── pkg
│       ├── config
│       │   └── config.go
│       ├── custom_api
│       │   ├── controllers
│       │   │   ├── networks.go
│       │   │   └── zones.go
│       │   └── router
│       │       └── routes.go
│       ├── ns1_custom_client
│       │   ├── models
│       │   │   ├── networks.go
│       │   │   └── zones.go
│       │   ├── services
│       │   │   ├── client.go
│       │   │   ├── dns_networks.go
│       │   │   └── dns_zones.go
│       │   └── utils
│       │       └── requests.go
│       ├── ns1-sdk
│       │   ├── models
│       │   └── sdk.go
│       └── official_api
│           ├── controllers
│           │   └── zones.go
│           └── router
│               └── routes.go
├── pkg
│   ├── api
│   │   ├── models
│   │   │   └── route.go
│   │   └── router.go
│   └── errors
│       └── errorHTTPResponse.go
|── README.md
└── LICENSE
```

The top level directories **cmd, internal, pkg** are commonly found in other popular Go projects, as explained in [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
