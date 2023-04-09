FROM golang:1.19-bullseye

RUN apt-get update
RUN apt-get install -y \
  python3 \
  python3-setuptools \
  python3-pip

RUN pip3 install --upgrade pip
RUN pip3 install pipenv

RUN curl -fsSL get.docker.com | bash

RUN curl https://get.ignite.com/cli! | bash

WORKDIR /app

ENTRYPOINT [ "container/entrypoint.sh" ]
