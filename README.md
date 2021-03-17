# Go-Microservices

Simple Example of adopting Repository-Service pattern with MySQL database in Go

This repo provides a simple example of adopting repository-service pattern with CRUD operations on a MySQL database.

Simple High-level Architecture

![](https://github.com/ShehabMMohamed/Go-Microservices/blob/main/High-Level-Architecture.png)

You can import `Go-Microservice-REST-API.postman_collection.json` for local testing.

#### Note 🔴: You need to install MySQL server and initialize a database ("todo") if you want to test it locally

### Dockerizing it

`docker build -t go-microservice .`

`docker run -p 8080:8080 go-microservice`

Give it a ⭐ if this repo was helpful to you
