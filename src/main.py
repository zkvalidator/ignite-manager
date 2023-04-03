import os
import sys
import subprocess
from pathlib import Path
import yaml
from jinja2 import FileSystemLoader, Environment

def load_config(config_file):
  with open(config_file, "r") as f:
    return yaml.safe_load(f)

def run_command(command):
  process = subprocess.Popen(command, shell=True)
  process.wait()
  if process.returncode != 0:
    print(f"Error: '{command}' failed with return code {process.returncode}")
    sys.exit(1)

def scaffold_chain(config):
  chain_name = config["chain"]["name"]
  chain_prefix = config["chain"]["prefix"]

  run_command(f"ignite scaffold chain {chain_name} --no-module --address-prefix {chain_prefix}")
  return chain_name

def scaffold_module(config, chain_name):
  module_name = config["module"]["name"]
  run_command(f"cd {chain_name} && ignite scaffold module --yes {module_name} --dep bank,staking")

def scaffold_models(config, chain_name):
  template_loader = FileSystemLoader(searchpath="./templates/")
  template_env = Environment(loader=template_loader)

  for model in config["models"]:
    model_type = model["type"]
    model_name = model["name"]
    model_attributes = " ".join(model["attributes"])

    run_command(f"cd {chain_name} && ignite scaffold {model_type} --yes --module {config['module']['name']} {model_name} {model_attributes}")

    if "events" in model:
      apply_event_template(config, model, chain_name, template_env)

def apply_event_template(config, model, chain_name, template_env):
  target = f"{chain_name}/x/{config['module']['name']}/keeper/msg_server_{model['name']}.go"
  template = template_env.get_template("event.go")
  rendered = template.render(model=model)

  with open(target, "a") as target_file:
    target_file.write(rendered)

def main():
  config_file = "config.yml"
  if not os.path.exists(config_file):
    print(f"Error: Configuration file '{config_file}' not found.")
    sys.exit(1)

  config = load_config(config_file)

  chain_name = scaffold_chain(config)
  scaffold_module(config, chain_name)
  scaffold_models(config, chain_name)

if __name__ == "__main__":
  print('Scaffolding chain...')
  main()
