FROM golang:1.14
LABEL maintainer="Tech Thinker <https://techthinker.js.org>"

RUN apt-get update && \
  apt-get install -y \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg-agent

RUN curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update && apt-get install -y migrate
RUN apt-get install -y protobuf-compiler && apt-get install -y golang-goprotobuf-dev


WORKDIR /go/src/app

ENTRYPOINT ["docker/entrypoint.sh"]
