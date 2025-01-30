#!/bin/bash

COMPOSE_FILE="compose/docker-compose.yaml"

function up() {
    echo "Starting services..."
    docker compose -f "$COMPOSE_FILE" up -d
}

function down() {
    echo "Stopping services..."
    docker compose -f "$COMPOSE_FILE" down
}

function test() {
    echo "Running tests..."
    curl -i http://localhost:3000/
    curl -i http://localhost:3000/backend
}

function clean() {
    echo "Cleaning up..."
    docker image rm my-backend
    docker image rm my-frontend
}

if [ $# -eq 0 ]; then
    echo "No arguments provided."
    echo "Usage: $0 {up|down|clean|test}"
    exit 1
fi

case $1 in
    up)
        up
        ;;
    down)
        down
        ;;
    clean)
        clean
        ;;
    test)
        test
        ;;
    *)
        echo "Invalid argument: $1"
        echo "Usage: $0 {up|down|clean|test}"
        exit 1
        ;;
esac
