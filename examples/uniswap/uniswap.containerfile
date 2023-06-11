FROM node:12-bullseye

WORKDIR /app

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list

RUN apt update
RUN apt install -y \
  apache2 \
  build-essential \
  curl \
  git \
  yarn

RUN npm install -g \
  fs \
  truffle-hdwallet-provider \
  web3

RUN git clone https://github.com/Uniswap/uniswap-interface.git
WORKDIR /app/uniswap-interface
RUN git clone https://github.com/second-state/how_to_deploy_uniswap.git
WORKDIR /app/uniswap-interface/how_to_deploy_uniswap/uniswap_v2
RUN node deploy_uniswap_v2.js

WORKDIR /app/uniswap-interface/how_to_deploy_uniswap/uniswap_interface
RUN python3 modify_addresses.py
WORKDIR /app/uniswap-interface
RUN yarn


