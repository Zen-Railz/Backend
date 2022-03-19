[![Go Report Card](https://goreportcard.com/badge/github.com/Zen-Railz/Backend)](https://goreportcard.com/report/github.com/Zen-Railz/Backend) 
[![CircleCI](https://circleci.com/gh/Zen-Railz/Backend/tree/main.svg?style=shield)](https://circleci.com/gh/Zen-Railz/Backend/tree/main)

# Overview
This is the backend repository of our application.
|||
|-|-|
| Repository | [GitHub](https://github.com/Zen-Railz/Backend) |
| Continuous Integration | [CircleCI](https://app.circleci.com/pipelines/github/Zen-Railz/Backend) |
| Hosting | [Heroku with Docker](https://devcenter.heroku.com/articles/container-registry-and-runtime) |
| Language | [Go](https://go.dev/) |
| Framework | [Gin](https://github.com/gin-gonic/gin) |
| Formatter | [gofmt](https://pkg.go.dev/cmd/gofmt) |
| Test | [Gingko](https://github.com/onsi/ginkgo) <br> [Gomega](https://github.com/onsi/gomega) <br> [gotestsum](https://github.com/gotestyourself/gotestsum) |

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