# Auth Service
This repository serves as a PoC for how the authentication may look for the Bee Well application.
## Tech
The server is written in Golang and utilises the `Gin` framework for handling HTTP requests. This application has been set up with a simple MQ pub/sub package that can be used for intercommunication between different services on Heroku. 
## Run Dev
Make sure to run a development PostgreSQL through docker using the following commands:
```
docker run -e POSTGRES_PASSWORD=password -p 5432:5432 postgres:alpine
docker run -d -h bee-well -p 5672:5672 rabbitmq:alpine
```