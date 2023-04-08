import os
import sys
import subprocess
from pathlib import Path
import yaml
from jinja2 import FileSystemLoader, Environment
import logging

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

  logging.debug(f"Loading configuration from '{config_file}'...")
  with open(config_file, "r") as f:
    config = yaml.load(f, Loader=Loader)
    logging.debug(f"Configuration loaded: {config}")
    return config

def run_command(command):
  process = subprocess.Popen(command, shell=True)
  process.wait()
  if process.returncode != 0:
    logging.debug(f"Error: '{command}' failed with return code {process.returncode}")
    sys.exit(1)

def scaffold_chain(config):
  chain_name = config["chain"]["name"]
  chain_prefix = config["chain"]["prefix"]
  logging.debug(f"Scaffolding chain '{chain_name}'...")
  run_command(f"ignite scaffold chain {chain_name} --no-module --address-prefix {chain_prefix}")
  return chain_name

def scaffold_module(config, chain_name):
  module_name = config["module"]["name"]
  logging.debug(f"Scaffolding module '{module_name}'...")
  run_command(f"cd {chain_name} && ignite scaffold module --yes {module_name} --dep bank,staking")

def scaffold_models(config, chain_name):
  template_loader = FileSystemLoader(searchpath="./templates/")
  template_env = Environment(loader=template_loader)

  for model in config["models"]:
    model_type = model["type"]
    model_name = model["name"]
    model_attributes = " ".join(model["attributes"])

    logging.debug(f"Scaffolding model '{model_name}'...")
    run_command(f"cd {chain_name} && ignite scaffold {model_type} --yes --module {config['module']['name']} {model_name} {model_attributes}")

    if "events" in model and model["events"] == True:
      apply_event_template(config, model, chain_name, template_env)

def update_go_mod(chain_name):
  replacements = [
    ("github.com/cosmos/cosmos-sdk v0.46.3", "github.com/cosmos/cosmos-sdk v0.46.2"),
    ("github.com/ignite/cli v0.25.1", "github.com/ignite/cli v0.25.0"),
  ]

  go_mod_path = f"{chain_name}/go.mod"
  logging.debug(f"Updating go.mod: {go_mod_path}...")

  with open(go_mod_path, "r") as go_mod_file:
    go_mod_content = go_mod_file.read()

  for search, replace in replacements:
    go_mod_content = go_mod_content.replace(search, replace)

  with open(go_mod_path, "w") as go_mod_file:
    go_mod_file.write(go_mod_content)

  logging.debug(f"Running go mod tidy: {chain_name}...")
  run_command(f"cd {chain_name} && go mod tidy")

def move_and_replace_config(chain_name, config):
  logging.debug(f"Moving and replacing config: {chain_name}/config.yml")
  os.remove(f"{chain_name}/config.yml")
  with open(f"{chain_name}/config.yml", "w") as f:
    yaml.dump(config["ignite"]["config"], f)

def run_ignite_commands(chain_name):
  logging.debug(f"Running ignite commands: {chain_name}...")
  run_command(f"cd {chain_name} && ignite chain build")

def apply_event_template(config, model, chain_name, template_env):
  target = f"{chain_name}/x/{config['module']['name']}/keeper/msg_server_{model['name']}.go"
  template = template_env.get_template("src/templates/event.go")
  rendered = template.render(model=model)

  with open(target, "a") as target_file:
    target_file.write(rendered)

def main():
  config_file = "build.yml"
  if not os.path.exists(config_file):
    logging.debug(f"Error: Configuration file '{config_file}' not found.")
    sys.exit(1)

  config = load_config(config_file)

  chain_name = scaffold_chain(config)
  move_and_replace_config(chain_name, config)
  update_go_mod(chain_name)
  scaffold_module(config, chain_name)
  scaffold_models(config, chain_name)
  run_ignite_commands(chain_name)

if __name__ == "__main__":
  logging.debug('Starting...')
  main()
  logging.debug('Done!')
