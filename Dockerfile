FROM golang:1.20.5 AS builder

#ENV GOPROXY=https://goproxy.cn

COPY . /src
WORKDIR /src

RUN make build

FROM ubuntu:20.04

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates netbase && \
    rm -rf /var/lib/apt/lists/ && \
    apt-get autoremove -y && \
    apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 6060
EXPOSE 6160
COPY configs /data/conf

CMD ["sh", "-c", "./server -c /data/conf"]
# enable k8s config map
#CMD ["sh", "-c", "./server -c /data/conf -n cinch -l layout"]
