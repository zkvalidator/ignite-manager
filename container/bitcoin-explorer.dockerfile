FROM node:20

WORKDIR /app

RUN git clone https://github.com/janoside/btc-rpc-explorer.git
WORKDIR /app/btc-rpc-explorer
RUN npm install

CMD [ "npm", "run", "start" ]
