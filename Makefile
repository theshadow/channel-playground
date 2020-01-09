.PHONY: all build run

STATSD_HOST=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' graphite)

all: build run

build:
	docker build --tag theshadow/channel-playground -f Dockerfile .

run:
	docker run \
      --env STATSD_INTERFACE=$(STATSD_HOST):8125 \
      theshadow/channel-playground:latest
