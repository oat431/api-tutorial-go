# API tutorial using GO lang and Gin framework

- just a simple tutorial to learn how to use Gin framework to build a RESTful API using GO lang

## Chapter list

1. Chapter 1: Project Setup
2. Chapter 2: Basic API
3. Chapter 3: Playing with Database
4. Chapter 4: CORS
5. Chapter 5: Pagination
6. Chapter 6: GraphQL
7. Chapter 7: Testing

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

### project structure

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

## Chapter 3: Playing with Database

In this chapter just create a basic API with gin-gonic associated with a database via GORM lib

you can check the code in the branch `chapter-3-playing-w-db`

on this chapter we create the following endpoints:

### Endpoints

- POST: `api/v2/users/` : create a new user
- GET : `api/v2/users/` : get all the users
- GET : `api/v2/users/<id>` : get a specific user
- PUT : `api/v2/users/<id>` : update a specific user
- DELETE : `api/v2/users/<id>` : delete a specific user

### project structure

``` bash
api-tutorial-go/
├─ configs/
│  ├─ db_config.go
│
├─ controllers/
│  ├─ user_v2_controller.go
│  
├─ dao/
│  ├─ user_dao.go
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
├─ repository/
│  ├─ user_repository.go
│
├─ routes/
│  ├─ route.go
│  
├─ services/
│  ├─ user_v2_service.go
│  
├─ utils/
│  ├─ dto-helper.go
│  ├─ mock-users.go

```

### Simple API Implementation with DB

1. Create a model
2. Create a DTO (Data Transfer Object)
3. Create a Request
4. Create a Repository (ORM or raw SQL)
5. Create a DAO (Data Access Object)
6. Create a Service (Business Logic)
7. Create a Controller and link it  with the Service
8. Link the controller with the routes

## Chapter 4: CORS

This chapter we need to make sure that our API is compatible with Frontend apps that why we need to enable CORS and to do that we need to install `cors` lib

### Focus on

``` bash
api-tutorial-go/
├─ configs/
│  ├─ cors_config.go
│
├─ routes/
│  ├─ route.go
│  
```

### To install it run the following command

``` bash
go get github.com/gin-contrib/cors
```

then create `cors_config.go` file add the following code:

``` go
package configs

import (
    "github.com/gin-contrib/cors"
)

func SetupCORS() cors.Config {
    return cors.Config{
        AllowOrigins:     []string{"http://url:port"},
        AllowMethods:     []string{"HTTP_METHOD_1", "HTTP_METHOD_2", "HTTP_METHOD_3" },
        AllowHeaders:     []string{"HEADER_NAME_1", "HEADER_NAME_2", "HEADER_NAME_3"},
        ExposeHeaders:    []string{"HEADER_NAME_1", "HEADER_NAME_2", "HEADER_NAME_3"},
        AllowCredentials: true,
    }
}
```

then in `route.go` file add the following code:

``` go
func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.Use(cors.New(configs.SetupCORS()))
    .
    .
}
```

** mark that you need to enable CORS before setting up the routes

## Chapter 5: Pagination

This chapter we need to add pagination to our API to make it more efficient
normally go lang community implement their own pagination but in this tutorial we will use `morkid/paginate` pagination

### Focus on

``` bash
api-tutorial-go/
├─ controllers/
│  ├─ user_v2_controller.go
│  
├─ dao/
│  ├─ user_dao.go
│
├─ payload/
│  ├─ response/
│  │  ├─ page_user_dto.go
│  
├─ repository/
│  ├─ user_repository.go
│
├─ routes/
│  ├─ route.go
│  
├─ services/
│  ├─ user_v2_service.go
│  
├─ utils/
│  ├─ dto-helper.go
```

### To install it run the following command

``` bash
go get github.com/morkid/paginate
```

then learning the `focus` on files
