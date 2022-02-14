# os-api

Simple Online Store API

## Prerequisite

* [Golang](https://golang.org) - Go Programming Language
* [Echo](https://echo.labstack.com/) - HTTP Framework
* [Gomod](https://github.com/golang/go/wiki/Modules) - Modules Management

## Install Modules

Run `go mod tidy` to install dependencies required that state on go.mod

Make sure you have go v11 or above installed

## Install Golang Migrate

[Please read here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

## Run Migration

Run `make migrate` or `go run main.go migrate`

## Development server

Run `make run` or `go run main.go` for a dev server. Navigate to `http://localhost:7723/`.

## Run Test

Run `make test` or `go clean -testcache && go test ./... | grep -v 'no test files'`.

## Build

Run `make build` or `go build -o os-api`.


## Case

- order cancelled due to mismatch (negative) qty

## Proposed Solution

- use unsigned data type for qty

## Author

* **Andhika Gama** - *Initial work* - [AndhikaGama](https://github.com/andhikagama)

