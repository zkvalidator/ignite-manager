# Ignite Scaffolding Tool

This project provides a simple and extensible Python-based tool to automate the scaffolding process for Tendermint and Cosmos SDK chains using the Ignite framework. The tool reads a YAML configuration file to scaffold the chain, module, and models, and it supports event templates and custom .go files.

## Features

- Chain and module scaffolding based on YAML configuration
- Model scaffolding with support for event templates
- Custom .go file integration
- Easy-to-use CLI

## Prerequisites

- Python 3.6 or later
- Ignite CLI installed and configured

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/ignite-scaffolding-tool.git
cd ignite-scaffolding-tool
```

2. (Optional) Set up a virtual environment to isolate the dependencies:

```bash
python -m venv venv
source venv/bin/activate  # For Windows, use `venv\Scripts\activate`
```

3. Install the required dependencies:

```bash
pip install -r requirements.txt
```

## Usage

1. Edit the `config.yml` file to define your chain, module, and models. You can also specify event templates and custom .go files.

The `config.yml` file is used to define your chain, module, and models, as well as specify event templates and custom .go files. The file is structured into sections that provide options for customization.

Here's a generic example of the `config.yml` file:

```yaml
chain:
  name: mychain
  prefix: myp

module:
  name: mymodule

models:
  - type: list
    name: entity
    attributes:
      - field1:string
      - field2:int
    events: true
    custom_files:
      - custom_entity.go

  - type: list
    name: resource
    attributes:
      - owner:string
      - name:string
      - category:string
      - value:int

# Add more models here
```

### YAML Options

#### Chain

- `name` (string): The name of your chain. This name will be used by the Ignite CLI when scaffolding the chain.
- `prefix` (string): The address prefix for your chain.

#### Module

- `name` (string): The name of your module. This name will be used by the Ignite CLI when scaffolding the module.

#### Models

A list of model objects, where each model object has the following keys:

- `type` (string): The type of the model. Supported types are `list`, `map`, and `single`.
- `name` (string): The name of the model.
- `attributes` (list of strings): A list of attributes for the model in the format `name:type`. Supported types include `string`, `int`, `bool`, `array`, and others.
- `events` (boolean, optional): If `true`, the scaffolding tool will apply the event template to the model. Default is `false`.
- `custom_files` (list of strings, optional): A list of custom .go files to be integrated into the scaffolded chain. The files should be placed in the project folder.

### Adding More Models

You can add more models to the `models` list in the `config.yml` file. Simply follow the structure of the existing model objects and provide the necessary keys and values for each new model.

Remember to update the `scaffold.py` tool and create any required templates or custom .go files as needed to support the new models.

2. Run the scaffolding tool:

```bash
python scaffold.py
```

The tool will scaffold the chain, module, and models according to the configuration in `config.yml`. It will also apply event templates and integrate custom .go files as specified.

## Customization

To customize the event templates, edit the `templates/event_template.go` file. You can use Jinja2 template syntax to dynamically generate the contents based on the model configuration.

To add custom .go files, create the files in your project folder and list them under the "custom_files" key in the `config.yml` file for the corresponding model. The tool will copy the custom .go files to the appropriate locations in the scaffolded chain.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.