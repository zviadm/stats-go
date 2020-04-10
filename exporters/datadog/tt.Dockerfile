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
# DataDog - DogStatsD
RUN apt-get install -y --no-install-recommends \
		apt-transport-https \
		gnupg \
	&& sh -c "echo 'deb https://apt.datadoghq.com/ stable 7' > /etc/apt/sources.list.d/datadog.list" \
	&& apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 A2923DFF56EDA6E76E55E492D3A80E30382E94DE \
	&& apt-get update && apt-get install -y --no-install-recommends datadog-dogstatsd

# Extra tools.
# ...

# Environment variables.
ENV GOPATH=/root/.cache/goroot:/root/src
ENV PATH=$PATH:/root/.cache/goroot/bin
