# GDSC_Backend

## Homesei

An android application using AR technology to create a safe living environment for the aging.

The Backend part for Homesei application.

## Competition
Homesei is created for [Google Solution Challenge 2023](https://developers.google.com/community/gdsc-solution-challenge).

## Homesei: The 17 Sustainable Development Goals of the United Nations
1. For [Good Health and Wellbeing](https://sdg-tracker.org/good-health).
2. For [Sustainable Cities and Communities](https://sdg-tracker.org/cities).


## Document & Record
[Document Record](https://xiaogeamadeus.notion.site/GDSC-Project-Homesei-b49e13ca38784a16b2c54a22cf8b1357)

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
