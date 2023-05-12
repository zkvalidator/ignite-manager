FROM golang:1.19-bullseye

RUN apt-get update
RUN apt-get install -y \
  jq \
  python3 \
  python3-setuptools \
  python3-pip

RUN pip3 install --upgrade pip
RUN pip3 install pipenv

RUN curl -fsSL get.docker.com | bash

WORKDIR /app
# RUN git clone https://github.com/ignite/cli ignite-cli --depth=1
RUN git clone https://github.com/ignite/cli ignite-cli
WORKDIR /app/ignite-cli
RUN git fetch
ARG config_ignite_version
ENV IGNITE_VERSION=${config_ignite_version}
RUN git checkout ${IGNITE_VERSION}
RUN make install
RUN cp /go/bin/ignite /usr/local/bin/ignite
# RUN env && exit 1
# RUN find / -name ignite && exit 1

# https://github.com/ignite/installer#usage
# RUN curl https://get.ignite.com/cli@${IGNITE_VERSION}! | bash

WORKDIR /app

ENTRYPOINT [ "container/ignite-manager.entrypoint.sh" ]
