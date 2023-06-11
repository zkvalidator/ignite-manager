FROM node:16

WORKDIR /app

RUN git clone https://github.com/xops/expedition
WORKDIR /app/expedition
RUN npm install

CMD [ "npm", "run", "start" ]
