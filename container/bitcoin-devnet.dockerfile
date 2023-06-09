# FROM debian:bullseye
FROM lightninglabs/bitcoin-core:latest

# RUN apt-get update
# RUN apt-get upgrade -y

# RUN apt-get install -y \
#   build-essential \
#   clang\
#   curl \
#   git \
#   jq \
#   libssl-dev \
#   make \
#   ncdu \
#   npm \
#   pkg-config \
#   snapd \
#   tar \
#   wget

# RUN apt-get install -y snapd
# ENV PATH="$PATH:/snap/bin"

# RUN snap install bitcoin-core

COPY bitcoin-devnet.entrypoint.sh /bitcoin-devnet.entrypoint.sh
RUN chmod +x /bitcoin-devnet.entrypoint.sh

ENTRYPOINT [ "/bitcoin-devnet.entrypoint.sh" ]
