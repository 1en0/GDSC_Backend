# GDSC_Backend

An android application using AR technology to create a safe living environment for the aging.

## Get Started
1. Please download the docker desktop in your pc by the link below.
   https://docs.docker.com/desktop/install/mac-install/
1. Docker compose up

   `docker compose up`

    If the database cannot be constructed, please use:

    `docker compose up --force-recreate --build`

## Build with docker

1. Build image

   `docker image build -t <container name> .`

1. Run container

   `docker container run <container name>`


Then, you can test the endpoint by port: http://localhost:8080/
