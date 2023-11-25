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