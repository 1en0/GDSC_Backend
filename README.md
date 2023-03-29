# GDSC_Backend

An android application using AR technology to create a safe living environment for the aging.

## Run development server with docker

1. Docker compose up

   `docker compose up`

## Build with docker

1. Build image

   `docker image build -t <container name> .`

1. Run container

   `docker container run <container name>`

## Get Start

1.  Download all necessary dependencies.

    `go mod tidy `

1.  Install necessary DB.

    `brew install mysql`

1.  Create local database `gdsc`

1.  Execute the [init.sql](config%2Finit.sql) file to construct table

    `go run init.sql`

1.  Start server

    `go run main.go`

Then, you can test the endpoint by port: http://localhost:8080/
