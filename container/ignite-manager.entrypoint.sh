#!/bin/bash

export PIPENV_VENV_IN_PROJECT=true

echo "received arguments: $@"

echo "installing depedencies..."
pipenv install

echo "building..."
pipenv run python3 src/main.py -c "${config_path}" $@
