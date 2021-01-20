# Todo App Project Backend

Backend App Api

## Requirements
[go](https://golang.org/doc/install)

[pact](https://github.com/pact-foundation/jest-pact)


## Install, Build And Tests



Compile development
```
$ go build
```

Lints and fixes files
```
$ golangci-lint run -c .config/.golangci.yaml -v
```

### Test Process

Run the tests
```
$ go test ./... (For all tests)
```

Run the provider test
```
$ CGO_ENABLED=0 go test -v test/pact/provider_test.go
```


## Directory Structure

```
.
├── .cd                   
│   ├── gitlab-ci.yml         - Using for gitlab pipeline configurations.
├── .config                
│   ├── .golangci.yaml        - Using for linter configurations.
│   ├── dev.yaml              - Port and appname conf  
│   ├── prod.yaml             - Port and appname conf       
├── .mocks         
│   ├── mock_todo_repository.go     - Generated file for mock repository.
│   ├── mock_todo_service.go        - Generated file for mock service.         
├── config
│   ├── config.go
│   ├── config_test.go              - Config test with TDD
├── handler
│   ├── handler.go   
│   ├── handler_test.go             - Handler test with TDD
├── model
│   ├── todo_model.go              
├── repository
│   ├── todo_repository.go
├── server
│   ├── server.go
├── service
│   ├── service.go
│   ├── service_test.go              - Service test with TDD
├── test
│   ├── pact
│     ├── log                  - This file has provider test logs
│     ├── provider_test.go     - This file is using for provider test
│   ├── testdata
│     ├── test-config.yaml              - This file is using for config port test          
├── Dockerfile         
├── k8s.yaml     - Kubernetes deployment yaml
├── main.go    
```

## Deployment Process
On Gitlab Pipeline
  - build
  - lint
  - test
  - pact-publish
  - package
  - deploy