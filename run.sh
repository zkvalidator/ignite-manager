#!/bin/bash

if [ -z "$1" ]; then
  $1 = "build.yml"
fi

COMPOSE_FILE="container/ignite-manager.docker-compose.yml"

export CONFIG_FILE="$1"

docker compose --file $COMPOSE_FILE up --build
docker compose --file $COMPOSE_FILE down
