# Auth Service
This repository serves as a PoC for how the authentication may look for the Bee Well application.
## Tech
The server is written in Golang and utilises the `Gin` framework for handling HTTP requests. This application has been set up with a simple MQ pub/sub package that can be used for intercommunication between different services on Heroku. 