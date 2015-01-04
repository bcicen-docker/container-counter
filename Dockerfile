#container-counter v0.1

FROM flynn/busybox
MAINTAINER Bradley Cicenas <bradley.cicenas@gmail.com>

ENV COUNTER_STATSD 127.0.0.1:8125

ADD build/container-counter.linux.64 /bin/container-counter

CMD ["/bin/container-counter"]
