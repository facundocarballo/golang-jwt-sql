# Golang - JWT - SQL

## Step by Step - Server Setup
> This dependecies are needed to connect our golang server with the MySQL database

1. Init Golang Project
```bash
go mod init github.com/facundocarballo/golang-mysql-connection
```
2. Init Github Repo
```bash
git init
```
3. Install GoDotEnv for .env files
```bash
    go get github.com/joho/godotenv
```
4. Install Go-SQL-Driver
```bash
    go get github.com/go-sql-driver/mysql
```
> It's important to know that you will have to import this library to connect your golang server with your mysql database. Normally using like this example.
```golang
    import (
        _ "github.com/go-sql-driver/mysql"
    )
```
5. Install JWT
```bash
    go get github.com/golang-jwt/jwt/v4
```
> This is used in the **crypto** module to generate and get the JWT.
6. Run the server
```bash
    go run main.go
```
---

## API Calls - Examples
### Login
#### POST REQUEST
As users we have to demostrate to the server that we are a valid user.
We have to make a POST REQUEST to this endpoint **'/login'** and sending in the body of the request this:
```json
    {
        "email": "some@email.com",
        "password": "strongPassword"
    }
```
> This data it's just an example, the email and passwords have to match with the data stored in the database.

This endpoint will return a JWT if all goes well.

Using curl will be something like this.
```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{
        "email": "some@email.com",
        "password": "strongPassword"
    }' \
    http://localhost:3690/login
```

### Task
This endpoint works for create a new task, and also for getting all the task associate with a particular user.
#### POST REQUEST
This request it's responsible of try to generate a new task for the user. The user that want's to create a new task will have to call to this endpoint **'/task'** using a POST REQUEST.
The body of the request will have to have this information
```json
    {
        "name": "task name",
        "description": "task description"
    }
```
Also we have to provide in the request our JWT.
Using curl will be something like this.
```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Mn0.wcx99moQmQ6n7UVjEU1Y6Jz2HPaTrOmffMQycsawnMo" \
    -d '{
        "name": "task name",
        "description": "task description"
    }' \
    http://localhost:3690/task
```
> This is a JWT of example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Mn0.wcx99moQmQ6n7UVjEU1Y6Jz2HPaTrOmffMQycsawnMo

#### GET REQUEST
This request it's responsible for get all the task associated to a particular user. That user is the owner of the JWT.
The user that wants to get all of your tasks, will have to call to this endpoint **'/task'** using a GET REQUEST.
Using curl will be something like this.
```bash
    curl -X GET \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Mn0.wcx99moQmQ6n7UVjEU1Y6Jz2HPaTrOmffMQycsawnMo" \
    http://localhost:3690/task
```
> This is a JWT of example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Mn0.wcx99moQmQ6n7UVjEU1Y6Jz2HPaTrOmffMQycsawnMo

### User
This endpoint works for create a new user, and also for getting all the users in our database.
#### POST REQUEST
To create a new user, you have to call to this endpoint **'/user'** with a POST REQUEST and provide this data in the body of the request.
```json
    {
        "name": "Name of the user",
        "email": "email@user.com",
        "password": "strongPassword",
    }
```
Here you don't need to provide a JWT, because your user it's not in the database before your request.
The server here can return to you the JWT, but for this case it's not implemented in that way :(
Using curl will be something like this.
```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Name of the user",
        "email": "email@user.com",
        "password": "strongPassword"
    }' \
    http://localhost:3690/user
```
#### GET REQUEST
To get all the users that are stored in this database, you will have to call to this endpoint **'/user'** with a GET REQUEST.
Here you don't need to provide any JWT, because it's not implemented.

Using curl will be something like this.
```bash
    curl http://localhost:3690/user
```