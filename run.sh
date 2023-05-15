#!/bin/bash

if [ -z "$1" ]; then
  $1="build.yml"
fi

# https://stackoverflow.com/a/21189044
function parse_yaml {
   local prefix=$2
   local s='[[:space:]]*' w='[a-zA-Z0-9_]*' fs=$(echo @|tr @ '\034')
   sed -ne "s|^\($s\):|\1|" \
        -e "s|^\($s\)\($w\)$s:$s[\"']\(.*\)[\"']$s\$|\1$fs\2$fs\3|p" \
        -e "s|^\($s\)\($w\)$s:$s\(.*\)$s\$|\1$fs\2$fs\3|p"  $1 |
   awk -F$fs '{
      indent = length($1)/2;
      vname[indent] = $2;
      for (i in vname) {if (i > indent) {delete vname[i]}}
      if (length($3) > 0) {
         vn=""; for (i=0; i<indent; i++) {vn=(vn)(vname[i])("_")}
         printf("%s%s%s=\"%s\"\n", "'$prefix'",vn, $2, $3);
      }
   }'
}

export config_path="$1"
echo "config file: $config_path"

# parse_yaml $1 "config_"
eval $(parse_yaml $config_path "config_")

export config_ignite_version="${config_ignite_version}"
echo "ignite version: $config_ignite_version"
export arguments="${@:2}"
echo "arguments: $arguments"

COMPOSE_FILE="container/ignite-manager.docker-compose.yml"

docker network create ignite-manager || true
docker compose --file $COMPOSE_FILE up --build
docker compose --file $COMPOSE_FILE down
