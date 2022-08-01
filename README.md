# REST API - PRIVY BE TEST - Golang

A simple REST API for PRIVY BE TEST using [gin-gonic](https://github.com/gin-gonic/gin) for the web framework and [gorm](https://gorm.io) for the database ORM.

The project includes the following features:
* RESTful API endpoint
* Standard CRUD operations
* Data validation
* GEO IP lookup
* [JWT](https://jwt.io/)-based authentication
* Database migration
* Postman documentation and API testing

## Getting Started
Please follow [the instructions](https://golang.org/doc/install) to install Go on your local machine. This project requires **Go 1.18 or higher+**.

We also need to install [Docker](https://www.docker.com/get-started) if you want to try the project without setting your own environment.

After all done, run the following commands to start the project:

```bash
# Clone this repo
git clone https://gitlab.com/mr687/privy-be-test.git

# Go to the project directory
cd privy-be-test

# Install dependencies
go mod init

# Duplicate the .env.example file to .env and fill in your own values
cp .env.example .env

# Run postgreSQL database via Docker compose
docker compose up -d
# docker-compose up -d

# Start the rest server
go run main.go

# Server running default on http://localhost:8080
```

## API Documentation
The API documentation is hosted on [Postman Doc](https://documenter.getpostman.com/view/1838168/Uze1uio4).
