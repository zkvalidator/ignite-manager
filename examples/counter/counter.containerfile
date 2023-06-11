FROM debian:bullseye

RUN apt update
RUN apt install -y curl wget

RUN curl -fsSL get.docker.com | bash

ENV ver "1.19.1"
WORKDIR /app
RUN wget "https://golang.org/dl/go$ver.linux-arm64.tar.gz"
RUN rm -rf /usr/local/go
RUN tar -C /usr/local -xzf "go$ver.linux-arm64.tar.gz"
RUN rm "go$ver.linux-arm64.tar.gz"
RUN echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> /root/.bashrc

RUN curl -L https://foundry.paradigm.xyz/ | bash
RUN PATH=$PATH:/usr/local/go/bin:$HOME/go/bin /root/.foundry/bin/foundryup

RUN /root/.foundry/bin/forge init --no-commit bitcoin-ethermint-app
WORKDIR /app/bitcoin-ethermint-app

COPY Counter.sol /app/bitcoin-ethermint-app/src/Counter.sol

RUN rm /app/bitcoin-ethermint-app/test/Counter.t.sol
COPY Counter.T.sol /app/bitcoin-ethermint-app/test/Counter.T.sol

RUN /root/.foundry/bin/forge test -vv

COPY Counter.s.sol /app/bitcoin-ethermint-app/script/Counter.sol

RUN mkdir -p /app/front
COPY index.html /app/front/index.html

COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
