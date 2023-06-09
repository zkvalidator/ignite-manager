FROM golang:1.19.1-bullseye

WORKDIR /app
RUN git clone https://github.com/celestiaorg/ethermint.git
WORKDIR /app/ethermint
RUN git checkout bitcoin-da
RUN make install

RUN apt update
RUN apt install -y jq

RUN ./init.sh

COPY ethermint.entrypoint.sh /ethermint.entrypoint.sh
RUN chmod +x /ethermint.entrypoint.sh

ENTRYPOINT [ "/ethermint.entrypoint.sh" ]
