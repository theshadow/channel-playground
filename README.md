# channel-playground

Somewhere to test out various channel architectures. Uses statsd to report the metrics.

https://hub.docker.com/r/graphiteapp/graphite-statsd

```
docker run -d\
 --name graphite\
 --restart=always\
 -p 80:80\
 -p 2003-2004:2003-2004\
 -p 2023-2024:2023-2024\
 -p 8125:8125/udp\
 -p 8126:8126\
 graphiteapp/graphite-statsd




docker build \
  --tag "theshadow/channel-playground" \
  .

# Get the IP of the instance and run the benchmark suite
IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' graphite) docker run \
  --env STATSD_INTERFACE=$IP:8125 \
  theshadow/channel-playground:latest
```