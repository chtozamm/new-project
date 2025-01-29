#!/bin/bash

COMPOSE_FILE="compose/docker-compose.yaml"

function clean() {
    docker image rm my-backend
    docker image rm my-frontend
}

clean
