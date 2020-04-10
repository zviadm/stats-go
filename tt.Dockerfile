FROM ubuntu:18.04

# Compilers and Build tools
WORKDIR /root
RUN apt-get update && apt-get install -y --no-install-recommends \
		autoconf \
		automake \
		ca-certificates \
		g++ \
		gcc \
		git \
		libc6-dev \
		libtool \
		make \
		pkg-config \
		python

ADD https://dl.google.com/go/go1.14.linux-amd64.tar.gz ./
RUN tar -xvzf go1.14.linux-amd64.tar.gz \
	&& mv go go1.14 \
	&& rm go1.14.linux-amd64.tar.gz

# ThirdParty dependencies. Sort dependencies: Slowest->Fastest.
# ....

# Extra tools.
# ....

# Environment variables.
ENV GOPATH=/root/.cache/goroot:/root/src
ENV PATH=$PATH:/root/.cache/goroot/bin
