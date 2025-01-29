#!/bin/bash

COMPOSE_FILE="compose/docker-compose.yaml"

function up() {
    docker compose -f "$COMPOSE_FILE" up -d
}

up
