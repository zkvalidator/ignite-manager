#!/bin/bash

docker compose --file container/docker-compose.yml up --build
docker compose --file container/docker-compose.yml down
