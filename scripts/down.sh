#!/bin/bash

COMPOSE_FILE="compose/docker-compose.yaml"

function down() {
    docker compose -f "$COMPOSE_FILE" down
}

down
