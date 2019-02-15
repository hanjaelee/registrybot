FROM jayhanjaelee/scratch:latest
MAINTAINER Jay Lee <jay@hanjaelee.com>

ADD . /registrybot

WORKDIR /registrybot

ENTRYPOINT ["./bin/registrybot"]