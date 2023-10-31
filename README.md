# API tutorial using GO lang and Gin framework

- just a simple tutorial to learn how to use Gin framework to build a RESTful API using GO lang

## Chapter list

1. Chapter 1: Project Setup
2. Chapter 2: Basic API
3. Chapter 3: Playing with Database
4. Chapter 4: CORS
5. Chapter 5: GraphQL
6. Chapter 6: Testing

## Project
Todo List Project
- CRUD todolsit
- CRUD sub task

## Deploy site
- on render

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

## Chapter 2: Basic API

In this chapter just create a basic API with gin-gonic (mock the data for now)

you can check the code in the branch `chapter-2-basic-api`

on this chapter we create the following endpoints:

### Endpoints

- POST: `api/v1/users/` : create a new user
- GET : `api/v1/users/` : get all the users
- GET : `api/v1/users/<id>` : get a specific user
- PUT : `api/v1/users/<id>` : update a specific user
- DELETE : `api/v1/users/<id>` : delete a specific user

### project structure:

``` bash
api-tutorial-go/
├─ controllers/
│  ├─ user-controller.go
│  
├─ models/
│  ├─ user.go
│  
├─ payload/
│  ├─ request/
│  │  ├─ user-request.go
│  ├─ response/
│  │  ├─ user-dto.go
│  
├─ routes/
│  ├─ route.go
│  
├─ services/
│  ├─ user-service.go
│  
├─ utils/
│  ├─ dto-helper.go
│  ├─ mock-users.go

```

### Simple API Implementation

1. Create a model
2. Create a DTO (Data Transfer Object)
3. Create a Request
4. Create a Service (Business Logic)
5. Create a Controller and link it  with the Service
6. Link the controller with the routes

and all the rest is learning by this example
