# Go Service Template

This repository contains a Golang service template
to start development of a new backend service in Go according to current standards
and a strict separation of concerns.

## Run the application

Is as easy as executing the following commands:

```shell script
$ go mod install
$ make run
```

## How to use this

1. Clone the project to a new folder
2. Edit `./cmd/main.go` to indicate the correct *project name, description and author*
3. Remove `.git` folder if you need a clean history, otherwise it's going to contain an initial commit reference
4. Initialize a new git repository if you deleted the `.git` folder
5. Update the file `./config/service.env` to config your service:
    - Adjust `DATABASE` vars to point to the correct database
    - Set `SERVICE` vars to your needs

## Build and run the application

1. **Build**

```shell script
make build
make docker-build
```

2. **Run**

```shell script
make run
make docker-run
```

3. **Test**

```shell script
go test -v ./test/...
```

_______

## Generate Docs

```shell script
# Get swag
go install github.com/swaggo/swag/cmd/swag

# Generate docs
swag init --dir cmd/api --parseDependency --output docs
```
Run and go to **http://localhost:3000/api-docs/index.html**

You might need to tell your system the place where the `swag` binary is
stored in case your `$GOPATH` is not set to that location. On MacOS this
is mostly `export GOPATH=$HOME/go/bin`. Also include this folder into your
`PATH` with `export PATH="$GOPATH:$PATH"`. 

## Frameworks and tools

1. Golang >= 17.5
2. [`gin`](https://github.com/gin-gonic/gin) for REST APIs
3. [`gorm`](https://gorm.io) as database object relation model
4. [`viper`](https://github.com/spf13/viper) for `.env` file configuration
5. [`zap`](https://github.com/uber-go/zap) for logging
6. [`swag`](https://github.com/swaggo/swag) for OpenAPI doc generation 
