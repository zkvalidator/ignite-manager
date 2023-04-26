#!/bin/bash

if [ -z "$1" ]; then
  $1="build.yml"
fi

COMPOSE_FILE="container/ignite-manager.docker-compose.yml"

export CONFIG_FILE="$1"
echo "config file: $CONFIG_FILE"

docker network create ignite-manager || true
docker compose --file $COMPOSE_FILE up --build
docker compose --file $COMPOSE_FILE down
