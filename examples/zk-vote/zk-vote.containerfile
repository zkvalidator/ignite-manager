FROM node:20

RUN apt update
RUN apt install -y \
  git \
  curl \
  build-essential

WORKDIR /app

RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

ENV PATH="/root/.cargo/bin:${PATH}"

RUN git clone https://github.com/iden3/circom.git
WORKDIR /app/circom
RUN cargo build --release
RUN cargo install --path circom
RUN ln -s /app/circom/target/release/circom /bin/circom

WORKDIR /app
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
