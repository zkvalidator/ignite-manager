import os
import sys
import subprocess
from pathlib import Path
import logging
import argparse

import yaml
from caseconverter import pascalcase, camelcase

logging.basicConfig(level=logging.DEBUG)

def load_config(config_file):

  class Loader(yaml.SafeLoader):
    def __init__(self, stream):
      self._root = os.path.split(stream.name)[0]
      super(Loader, self).__init__(stream)
    def include(self, node):
      filename = os.path.join(self._root, self.construct_scalar(node))
      logging.debug(f"Loading included file: {filename}, root: {self._root}, node: {node}")
      with open(filename, 'r') as f:
        return yaml.load(f, Loader)
  Loader.add_constructor('!include', Loader.include)

  if not os.path.exists(config_file):
    logging.error(f"Error: Configuration file '{config_file}' not found.")
    sys.exit(1)
  logging.info(f"Loading configuration from '{config_file}'...")

  with open(config_file, "r") as f:
    config = yaml.load(f, Loader=Loader)
    logging.debug(f"Configuration loaded: {config}")
    return config

def run_command(command):
  process = subprocess.Popen(command, shell=True)
  process.wait()
  if process.returncode != 0:
    logging.error(f"Error: '{command}' failed with return code {process.returncode}")
    sys.exit(1)

def scaffold_chain(config):
  chain_name = config["chain"]["name"]
  chain_prefix = config["chain"]["prefix"]
  logging.info(f"Scaffolding chain '{chain_name}'...")
  run_command(f"ignite scaffold chain {chain_name} --no-module --address-prefix {chain_prefix}")
  # with open(f"{chain_name}/.gitignore", "w") as f:
  #   f.write("*\n")

def switch_framework(config, chain_name):
  if config["ignite"]["framework"]["type"] == "rollkit":
    logging.info(f"Switching to rollkit framework: {chain_name}...")
    run_command(f"""cd {chain_name} \
      && go mod edit -replace github.com/cosmos/cosmos-sdk=github.com/rollkit/cosmos-sdk@{config["ignite"]["framework"]["versions"]["cosmos-sdk"]} \
      && go mod edit -replace github.com/tendermint/tendermint=github.com/celestiaorg/tendermint@{config["ignite"]["framework"]["versions"]["tendermint"]} \
      && go mod tidy \
      && go mod download \
    """)

def scaffold_module(config, chain_name):
  module_name = config["module"]["name"]
  logging.info(f"Scaffolding module '{module_name}'...")
  run_command(f"cd {chain_name} && ignite scaffold module --yes {module_name} --dep bank,staking")

def scaffold_models(config, chain_name):

  for model in config["models"]:
    model_type = model["type"]
    model_name = model["name"]
    model_attributes = " ".join(model["attributes"])

    logging.info(f"Scaffolding model '{model_name}'...")
    run_command(f"cd {chain_name} && ignite scaffold {model_type} --yes --module {config['module']['name']} {model_name} {model_attributes}")

    if "events" in model and model["events"] == True:
      apply_event_template(config, model, chain_name)

def update_go_mod(chain_name):
  replacements = [
    ("github.com/cosmos/cosmos-sdk v0.46.3", "github.com/cosmos/cosmos-sdk v0.46.2"),
    ("github.com/ignite/cli v0.25.1", "github.com/ignite/cli v0.25.0"),
  ]

  go_mod_path = f"{chain_name}/go.mod"
  logging.info(f"Updating go.mod: {go_mod_path}...")

  with open(go_mod_path, "r") as go_mod_file:
    go_mod_content = go_mod_file.read()

  for search, replace in replacements:
    go_mod_content = go_mod_content.replace(search, replace)

  with open(go_mod_path, "w") as go_mod_file:
    go_mod_file.write(go_mod_content)

  logging.info(f"Running go mod tidy: {chain_name}...")
  run_command(f"cd {chain_name} && go mod tidy")

def move_and_replace_config(chain_name, config):
  logging.info(f"Moving and replacing config: {chain_name}/config.yml")
  os.remove(f"{chain_name}/config.yml")
  with open(f"{chain_name}/config.yml", "w") as f:
    yaml.dump(config["ignite"]["config"], f)

def start(config, chain_name):
  logging.info(f"Building chain: {chain_name}...")
  run_command(f"cd {chain_name} && ignite chain build")

  daemon = f"/root/go/bin/{chain_name}d"
  chain_id = chain_name
  validator_name = chain_name
  key_name = chain_name

  if options.erase:
    logging.info(f"Erasing chain data: {chain_name}...")
    if config["ignite"]["framework"]["type"] == "rollkit":
      run_command("docker compose --file container/celestia-devnet.docker-compose.yml down")
    run_command(f"rm -rf ~/.{chain_name}")
    if config["ignite"]["framework"]["type"] == "rollkit":
      run_command(f"scripts/init.sh {daemon} {chain_id} {validator_name} {key_name}")

  if options.start:
    if config["ignite"]["framework"]["type"] == "rollkit":
      logging.info("Starting Celestia node...")
      run_command("docker compose --file container/celestia-devnet.docker-compose.yml up -d")
      logging.info(f"Starting rollup: {chain_name}...")
      run_command(f"scripts/start.sh {daemon}")
  return

def apply_event_template(config, model, chain_name):

  target = f"{chain_name}/x/{config['module']['name']}/keeper/msg_server_{model['name']}.go"
  search = f"""
	id := k.Append{pascalcase(model['name'])}(
		ctx,
		{camelcase(model['name'])},
	)"""
  insert = f"""
	ctx.EventManager().EmitTypedEvent(&types.EventCreate{pascalcase(model['name'])}{{
		Id: id,
	}})"""

  with open(target, "r") as target_file:
    content = target_file.read()
  content = content.replace(search, f'{search}\n{insert}')
  with open(target, "w") as target_file:
    target_file.write(content)

  append = f"""
message EventCreate{pascalcase(model['name'])} {{
	uint64 id = 1;
}}
"""

  target = f"{chain_name}/proto/{config['chain']['name']}/{config['module']['name']}/{model['name']}.proto"
  with open(target, "a") as target_file:
    target_file.write(append)

options = None

def main():
  parser = argparse.ArgumentParser(description="Ignite Manager Script")

  parser.add_argument(
    "-c",
    "--config",
    help="Path to the configuration file",
    default="build.yml",
  )

  parser.add_argument(
    "-s",
    "--start",
    help="Start the chain",
    default=True,
  )

  # erase data
  parser.add_argument(
    "-e",
    "--erase",
    help="Erase the chain data",
    default=True,
  )

  options = parser.parse_args()
  config_file = options.config
  config = load_config(config_file)

  chain_name = config["chain"]["name"]

  if os.path.exists(chain_name):
    logging.warn(f"Removing old chain: {chain_name}...")
    run_command(f"rm -rf {chain_name}")

  scaffold_chain(config)
  switch_framework(config, chain_name)
  move_and_replace_config(chain_name, config)
  # update_go_mod(chain_name)
  scaffold_module(config, chain_name)

  scaffold_models(config, chain_name)
  start(config, chain_name)

if __name__ == "__main__":
  logging.info('Starting...')
  main()
  logging.info('Done!')
