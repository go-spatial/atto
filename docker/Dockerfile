FROM ubuntu:latest
CMD bash

RUN apt-get update && apt-get -y install \
  apt-utils \
  build-essential \
  curl \
  git \
  tar \
  wget \
  zlib1g-dev \
  automake \
  libtool \
  xutils-dev \
  make \
  pkg-config \
  python-pip \
  libcurl4-openssl-dev \
  libpng-dev \
  libsqlite3-dev \
  cmake \
  ninja-build \
  clang \
  nodejs \
  npm \
  libxi-dev \
  libglu1-mesa-dev \
  x11proto-randr-dev \
  x11proto-xext-dev \
  libxrandr-dev \
  x11proto-xf86vidmode-dev \
  libxxf86vm-dev \
  libxcursor-dev \
  libxinerama-dev

RUN curl -o /tmp/go.tar.gz https://storage.googleapis.com/golang/go1.10.1.linux-amd64.tar.gz
RUN tar -xf /tmp/go.tar.gz -C /usr/local/bin
ENV GOPATH=/opt/

RUN mkdir -p /opt/src/github.com/mapbox
RUN git -C /opt/src/github.com/mapbox clone https://github.com/mapbox/mapbox-gl-native.git
RUN cd /opt/src/github.com/mapbox/mapbox-gl-native && \
	make linux-core

ENV CFLAGS="-I/opt/src/github.com/mapbox/mapbox-gl-native/include -I/opt/src/github.com/mapbox/mapbox-gl-native/platform/default"


RUN git clone https://github.com/mapbox/mason-js.git && \
	cd mason-js && \
	npm install && \
	npm link


ADD . /opt/src/github.com/go-spatial/atto
RUN cd /opt/src/github.com/go-spatial/atto && make

