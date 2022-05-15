[![Go Report Card](https://goreportcard.com/badge/github.com/Zen-Railz/Backend)](https://goreportcard.com/report/github.com/Zen-Railz/Backend) 
[![CircleCI](https://circleci.com/gh/Zen-Railz/Backend/tree/main.svg?style=shield)](https://circleci.com/gh/Zen-Railz/Backend/tree/main)

![ZenRailz Backend Logo](image/logo.svg)

# Overview
This is the backend repository of our application.
| Domain | Description |
|-|-|
| Repository | [GitHub](https://github.com/Zen-Railz/Backend) |
| Language | [Go](https://go.dev/) |
| Framework | [Gin](https://github.com/gin-gonic/gin) |
| Hosting | [Heroku with Docker](https://devcenter.heroku.com/articles/container-registry-and-runtime) |
| Database | [Heroku Postgres](https://www.heroku.com/postgres) |
| Continuous Integration | [CircleCI](https://app.circleci.com/pipelines/github/Zen-Railz/Backend) |
| Test | [Gingko](https://github.com/onsi/ginkgo) <br> [Gomega](https://github.com/onsi/gomega) <br> [gotestsum](https://github.com/gotestyourself/gotestsum) |
| Formatter | [gofmt](https://pkg.go.dev/cmd/gofmt) |
| Dependency Injection | [Wire](https://github.com/google/wire) |

# Preparation
## Installation of dependencies
```cmd
go mod download
```

# Development
The development of this project is on a windows machine. Steps described in the following sections will be based on windows commands.

### Setup - Local Environment
1. Ensure that port number is set in the environment variable. The local server will be listening on this port.
```cmd
SET PORT=3000
```
2. Ensure that database url is set in the environment variable.
```cmd
SET DATABASE_URL=<uri>
```
3. Run the project from the root directory.
```cmd
go run main.go
```

### Dependency Injection
As dependency injection is one of the most important design principles to reduce tight-coupling among components, we made use of Google's code generation tool, Wire, to achieve this.

To generate the dependency graph, run in the root directory
```cmd
wire ./...
```

# Test
## Testing Locally
As we are using the Ginkgo testing framework, we can run all the tests in this module with any of the below commands in the root folder.
```cmd
ginkgo ./...
```
or
```cmd
go test -v ./...
```
To view test files, turn on the visibility of `_test.go` files in the `.vscode/settings.json` file by setting the config to false, or simply comment it.
```json
{
  "files.exclude": {
    // "**/*_test.go": true
  }
}
```

# Deployment
1. Create a docker file
2. Create a circleci yml file
3. Push code to code repository