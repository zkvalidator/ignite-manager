# Ignite Manager

Ignite Manager is a project that automates the process of scaffolding, building, and configuring blockchain applications using the Ignite framework. It reads a configuration file and generates the necessary files and structure for your blockchain application.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
  - [build.yml](#buildyml)
  - [config.yml](#configyml)
- [Usage](#usage)
- [License](#license)

## Requirements

- Docker

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/zkvalidator/ignite-manager.git
   ```

2. Change the current directory to the project root:

   ```
   cd ignite-manager
   ```

3. Build and start the Docker image:

   ```
   ./run.sh build.yml
   ```

## Configuration

The project uses two main configuration files: `build.yml` and `config.yml`.

### build.yml

`build.yml` is the main configuration file that defines the structure of your blockchain application. It contains information about the chain, modules, and models. Here's an example of a `build.yml` file:

```yaml
ignite:
  config: !include config.yml
  framework:
    type: rollkit
    versions:
      github.com/cosmos/cosmos-sdk: github.com/rollkit/cosmos-sdk@v0.46.7-rollkit-v0.7.2-no-fraud-proofs
      github.com/tendermint/tendermint: github.com/celestiaorg/tendermint@v0.34.22-0.20221202214355-3605c597500d
    # type: cosmos-sdk
    # versions:
    #   github.com/cosmos/cosmos-sdk: github.com/cosmos/cosmos-sdk@v0.46.2
    #   github.com/ignite/cli: github.com/ignite/cli@v0.25.0

manager:
  start:
    chain_id: examplechain-0
    validator_name: examplevalidator
    key_name: examplevalidator
  tokens:
    - symbol: examplesym

chain:
  name: examplechain
  prefix: ex

modules:
  - name: examplemodule
    deps:
      - bank
      - staking

    models:
      - name: entity_name
        type: list
        attributes:
          - field1:string
          - field2:int
        events: true
        # TODO implement support for custom files
        # custom_files:
        #   - custom_entity.go

      - name: resource_name
        type: list
        attributes:
          - owner:string
          - name:string
          - category:string
          - value:int
```

Here's a brief explanation of the configuration options:

- `ignite`: Contains configurations related to the Ignite framework.
  - `config`: Includes the `config.yml` file.
  - `framework`: Specifies the framework type and versions.
    - `type`: The framework type, e.g., "rollkit".
    - `versions`: Specifies the versions of the required dependencies, e.g., "cosmos-sdk" and "tendermint".
- `chain`: Contains configurations related to the blockchain.
  - `name`: The name of the blockchain.
  - `prefix`: The address prefix for the blockchain.
- `module`: Contains configurations related to the module.
  - `name`: The name of the module.
- `models`: A list of models to be generated.
  - `name`: The name of the model.
  - `type`: The type of the model, e.g., "list".
  - `attributes`: A list of attributes for the model.
  - `events`: If set to `true`, events will be generated for the model.
  - `custom_files`: A list of custom files to be included in the model.

### config.yml

`config.yml` is [Ignite's configuration file (ref)](https://docs.ignite.com/references/config) that contains information about accounts, validators, and other configurations. Here's an example of a `config.yml` file:

```yaml
version: 1

accounts:
  - name: alice
    coins: ["20000token", "200000000stake"]
    mnemonic: winter blur imitate this open palace reward steel local noodle believe into evil other rebuild ready fuel someone body capital review mixture absurd seminar

validators:
  - name: alice
    bonded: "100000000stake"

init:
  config:
    api:
      enable: true
      swagger: false
      address: "tcp://0.0.0.0:1317"
      max-open-connections: 1000

client:
  openapi:
    path: "docs/static/openapi.yml"
  typescript:
    path: "ts-client"
  vuex:
    path: "vue/src/store"

faucet:
  name: alice
  coins: ["5token", "100000stake"]
```

Here's a brief explanation of the configuration options:

- `version`: The version of the configuration file.
- `accounts`: A list of accounts to be created.
  - `name`: The name of the account.
  - `coins`: A list of coins to be assigned to the account.
  - `mnemonic`: The mnemonic of the account.
- `validators`: A list of validators to be created.
  - `name`: The name of the validator.
  - `bonded`: The number of bonded tokens for the validator.
- `init`: Contains initialization configurations.
  - `config`: Contains API and client configurations.
    - `api`: Contains API configurations.
      - `enable`: If set to `true`, the API will be enabled.
      - `swagger`: If set to `true`, the Swagger UI will be enabled.
      - `address`: The address to bind the API.
      - `max-open-connections`: The maximum number of open connections for the API.
    - `client`: Contains client configurations.
      - `openapi`: Specifies the path for the OpenAPI documentation.
      - `typescript`: Specifies the path for the TypeScript client.
      - `vuex`: Specifies the path for the Vuex store.
- `faucet`: Contains faucet configurations.
  - `name`: The name of the account to be used as a faucet.
  - `coins`: A list of coins to be assigned to the faucet.

## Usage

1. Update the `build.yml` and `config.yml` files according to your desired blockchain structure and configurations.

2. Run the `run.sh` script to start the Ignite Manager:

   ```
   ./run.sh
   ```

3. The script will build your blockchain application based on the provided configurations.

## License

[MIT](LICENSE)