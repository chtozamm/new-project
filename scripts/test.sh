#!/bin/bash

COMPOSE_FILE="compose/docker-compose.yaml"

function test() {
    curl -i http://localhost:3000/
    curl -i http://localhost:3000/backend
}

test
