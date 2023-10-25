FROM golang:1.20.5 AS builder

#ENV GOPROXY=https://goproxy.cn
# CGO_ENABLED=1, need check `ldd --version` is same as builder
ENV CGO_ENABLED=0

# install cinch tool
RUN go install github.com/go-cinch/cinch/cmd/cinch@latest

WORKDIR /src
# download first can use docker build cache if go.mod not change
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates netbase && \
    rm -rf /var/lib/apt/lists/ && \
    apt-get autoremove -y && \
    apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /go/bin/cinch /go/bin

WORKDIR /app

COPY configs /data/conf

CMD ["sh", "-c", "./server -c /data/conf"]
