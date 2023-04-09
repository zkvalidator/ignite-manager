#!/bin/bash

if [ -z "$1" ]; then
  $1 = "build.yml"
fi

export CONFIG_FILE="$1"

docker compose --file container/docker-compose.yml up --build
docker compose --file container/docker-compose.yml down
