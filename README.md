# container-counter

container-counter is a small container that listens to docker events and increments a statsd gauge whenever a new container is created.

## Usage

Start a countainer-counter container with the local docker socket mounted as a volume, and provide a statsd instance via the env var COUNTER_STATSD:
```
docker run -d -e COUNTER_STATSD=1.2.3.4:8125 -v /var/run/docker.sock:/var/run/docker.sock bcicen/container-counter
```
That's it!

## TODO

Add counters for other pseudo-trivial docker metrics(number of images pulled, number of pushes performed, etc.)
