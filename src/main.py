import os
import sys
import subprocess
from pathlib import Path
import logging
import argparse
import json

import yaml
import toml
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

def run_command(command, input_data=None):
  process = subprocess.Popen(command, shell=True, stdin=subprocess.PIPE)
  if input_data is not None:
    process.communicate(input_data.encode())
  process.wait()
  if process.returncode != 0:
    logging.error(f"Error: '{command}' failed with return code {process.returncode}")
    sys.exit(1)

def scaffold_chain(config):

  chain_name = config["chain"]["name"]
  chain_prefix = config["chain"]["prefix"]

  if os.path.exists(f"build/{chain_name}"):
    logging.warning(f"Removing old chain: build/{chain_name}...")
    run_command(f"rm -rf build/{chain_name}")

  logging.info(f"Scaffolding chain '{chain_name}'...")
  run_command(f"cd build && ignite scaffold chain {chain_name} --no-module --address-prefix {chain_prefix}")

def update_go_mod(config, chain_name):
  if config["ignite"]["framework"]["versions"]:
    for key, value in config["ignite"]["framework"]["versions"].items():
      logging.info(f"Updating go.mod: {key}={value}")
      run_command(f"cd build/{chain_name} && go mod edit -replace {key}={value}")
    run_command(f"""cd build/{chain_name} \
      && go mod tidy \
      && go mod download \
    """)

def scaffold_modules(config, chain_name):
  for module in config["modules"]:
    module_name = module["name"]
    module_deps = ",".join(module["deps"])
    logging.info(f"Scaffolding module '{module_name}'...")
    run_command(f"cd build/{chain_name} && ignite scaffold module --yes {module_name} --dep {module_deps}")

    for model in module["models"]:
      model_type = model["type"]
      model_name = model["name"]
      model_attributes = " ".join(model["attributes"])

      logging.info(f"Scaffolding model '{model_name}' in module {module_name}...")
      run_command(f"cd build/{chain_name} && ignite scaffold {model_type} --yes --module {module_name} {model_name} {model_attributes}")

      if "events" in model and model["events"] == True:
        apply_event_template(module_name, model_name, chain_name)

def apply_event_template(module_name, model_name, chain_name):

  target = f"build/{chain_name}/x/{module_name}/keeper/msg_server_{model_name}.go"
  search = f"""
	id := k.Append{pascalcase(model_name)}(
		ctx,
		{camelcase(model_name)},
	)"""
  insert = f"""
	ctx.EventManager().EmitTypedEvent(&types.EventCreate{pascalcase(model_name)}{{
		Id: id,
	}})"""

  with open(target, "r") as target_file:
    content = target_file.read()
  content = content.replace(search, f'{search}\n{insert}')
  with open(target, "w") as target_file:
    target_file.write(content)

  append = f"""
message EventCreate{pascalcase(model_name)} {{
	uint64 id = 1;
}}
"""

  target = f"build/{chain_name}/proto/{chain_name}/{module_name}/{model_name}.proto"
  with open(target, "a") as target_file:
    target_file.write(append)

def move_and_replace_config(chain_name, config):
  logging.info(f"Moving and replacing config: build/{chain_name}/config.yml")
  os.remove(f"build/{chain_name}/config.yml")
  with open(f"build/{chain_name}/config.yml", "w") as f:
    yaml.dump(config["ignite"]["config"], f)

def build(chain_name):
  logging.info(f"Building chain: {chain_name}...")
  run_command(f"cd build/{chain_name} && ignite chain build")

def start_explorer(config, chain_name):
  chain = {
    "chain_name": chain_name,
    "coingecko": "",
    "api": [f"http://ignite-manager:1317"],
    "rpc": [f"http://ignite-manager:26657"],
    "sdk_version": "0.46.7",
    "coin_type": "118",
    "min_tx_fee": "800",
    "addr_prefix": config["chain"]["prefix"],
    "logo": "https://dl.airtable.com/.attachments/e54f814bba8c0f9af8a3056020210de0/2d1155fb/cosmos-hub.svg",
    "snapshot_provider": "",
    "assets": []
  }
  for token in config["manager"]["tokens"]:
    chain["assets"].append({
      "base": f"u{token['symbol']}",
      "symbol": token["symbol"].upper(),
      "exponent": "6",
      "coingecko_id": "",
      "logo": "https://dl.airtable.com/.attachments/e54f814bba8c0f9af8a3056020210de0/2d1155fb/cosmos-hub.svg"
    })

  with open(f"build/pingpub.json", "w") as f:
    json.dump(chain, f, indent=2)

  # run_command("pwd")
  # run_command("ls -la build/pingpub.json")
  run_command(f"docker compose --file container/pingpub.docker-compose.yml up --build -d")
  run_command(f"docker compose --file container/pingpub.docker-compose.yml logs -f --tail 100 &")

def deep_update(d, u):
  for k, v in u.items():
    if isinstance(v, dict):
      d[k] = deep_update(d.get(k, {}), v)
    else:
      d[k] = v
  return d

def start(config, chain_name):

  daemon = f"/go/bin/{chain_name}d"
  chain_id = config["manager"]["start"]["chain_id"]
  validator_name = config["manager"]["start"]["validator_name"]
  key_name = config["manager"]["start"]["key_name"]
  global options

  if options.erase:
    if config["ignite"]["framework"]["type"] == "rollkit":
      logging.info("Stopping Celestia node...")
      run_command("docker compose --file container/celestia-devnet.docker-compose.yml down || true")
    logging.info(f"Erasing chain data: {chain_name}...")
    run_command(f"rm -rf ~/.{chain_name}")

    # if config["ignite"]["framework"]["type"] == "rollkit":
    logging.info(f"Initializing chain: {chain_name}...")
    # run_command(f"src/scripts/init.sh {daemon} {chain_id} {validator_name} {key_name}")
    logging.debug("Resetting Tendermint...")
    run_command(f"{daemon} tendermint unsafe-reset-all")
    logging.debug("Initializing chain...")
    run_command(f"{daemon} init {validator_name} --chain-id {chain_id}")
    for key in config["ignite"]["config"]["accounts"]:
      key_name = key["name"]
      logging.debug(f"Adding key: {key_name}")
      if "mnemonic" in key:
        run_command(f"{daemon} keys add {key_name} --keyring-backend test --recover", input_data=key["mnemonic"])
      else:
        run_command(f"{daemon} keys add {key_name} --keyring-backend test")
      logging.debug(f"Adding genesis account: {key_name}")
      run_command(f"{daemon} add-genesis-account {key_name} {key['coins'][0]} --keyring-backend test")
      logging.debug(f"Generating gentx: {key_name}")
    for validator in config["ignite"]["config"]["validators"]:
      key_name = validator["name"]
      run_command(f"mkdir -p /root/.{chain_name}/config/gentx")
      run_command(f"{daemon} gentx {key_name} {validator['bonded']} --chain-id {chain_id} --keyring-backend test --output-document /root/.{chain_name}/config/gentx/gentx-{key_name}.json")
    logging.debug("Collecting gentxs...")
    run_command(f"{daemon} collect-gentxs")

  if "node_config" in config["manager"]:
    for filename in ["app", "client", "config"]:
      if filename in config["manager"]["node_config"]:
        logging.debug(f"Updating {filename}.toml")
        with open(f"/root/.{chain_name}/config/{filename}.toml", "r") as f:
          content = f.read()
        updated = deep_update(toml.loads(content), config["manager"]["node_config"][filename])
        with open(f"/root/.{chain_name}/config/{filename}.toml", "w") as f:
          f.write(toml.dumps(updated))

  if options.start:
    if config["ignite"]["framework"]["type"] == "rollkit":
      logging.info("Starting Celestia node...")
      run_command("docker compose --file container/celestia-devnet.docker-compose.yml up -d")
      run_command("docker compose --file container/celestia-devnet.docker-compose.yml logs -f --tail 100 &") 
      logging.info(f"Starting rollup: {chain_name}...")
      run_command(f"src/scripts/start.sh {daemon}")

def parse_args():
  parser = argparse.ArgumentParser(description="Ignite Manager Script")

  parser.add_argument(
    "-c",
    "--config",
    help="Path to the configuration file",
    default="build.yml",
  )

  parser.add_argument(
    "-r",
    "--rescaffold",
    help="Rescaffold the chain",
    default=False,
  )

  parser.add_argument(
    "-e",
    "--erase",
    help="Erase the chain data",
    default=True,
  )

  parser.add_argument(
    "-s",
    "--start",
    help="Start the chain",
    default=True,
  )

  parser.add_argument(
    "-x",
    "--explorer",
    help="Start the explorer",
    default=True,
  )

  options = parser.parse_args()
  return options

def main():
  global options
  options = parse_args()
  config_file = options.config
  config = load_config(config_file)

  chain_name = config["chain"]["name"]

  if options.rescaffold or not os.path.exists(f"build/{chain_name}"):
    scaffold_chain(config)
    update_go_mod(config, chain_name)
    move_and_replace_config(chain_name, config)
    scaffold_modules(config, chain_name)
    build(chain_name)

  if options.explorer:
    start_explorer(config, chain_name)

  start(config, chain_name)

if __name__ == "__main__":
  logging.info('Starting...')
  main()
  logging.info('Done!')
