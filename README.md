# API tutorial using GO lang and Gin framework

- just a simple tutorial to learn how to use Gin framework to build a RESTful API using GO lang

## Chapter list

1. Chapter 1: Project Setup
2. Chapter 2: Basic API
3. Chapter 3: Playing with Database
4. Chapter 4: CORS
5. Chapter 5: Testing
6. Chapter 6: Todo List Project

## Chapter 1: Project Setup

### Prerequisites

- [GO lang](https://golang.org/dl/)

### Setup

- clone the repo
- run `go get -u github.com/gin-gonic/gin`

### Run this project

- run `go run main.go`

### Check if the project is running

open your browser and go to `http://localhost:8080/api/v1/hello-world`
you should see a resposne like this:

``` json
{
    "message": "Hello World"
}
```
