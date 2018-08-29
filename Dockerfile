FROM golang

RUN apt-get update && \
    apt-get install -y \
      # build tools, for compiling
      build-essential \
      # install curl to fetch dev things
      curl \
      # we'll need git for fetching golang deps
      git

# install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# setup the app dir/working directory
RUN mkdir -p /go/src/github.com/tylerflint/survey
WORKDIR /go/src/github.com/tylerflint/survey
