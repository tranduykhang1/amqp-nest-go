#!/bin/bash

#Create networking
docker network create ammq-network

#Coppy env file
cp ./.env.example ./rabbitmq/.env
cp ./.env.example ./go-client/.env

#Start Nest-servier
cd rabbitmq
docker-compose up -d
cd ..



#Start Go-client
cd go-client
docker-compose up -d
