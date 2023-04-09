FROM node:16

RUN git clone https://github.com/ping-pub/explorer /app

WORKDIR /app
RUN yarn
COPY build/pingpub.json /app/src/chains/mainnet/test.json

CMD yarn serve
