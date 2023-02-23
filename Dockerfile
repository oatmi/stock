FROM --platform=linux/x86_64 golang:1.20.1

RUN apt-get update \
    && apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross git

COPY . ./src/github.com/oatmi/stock

RUN cd ./src/github.com/oatmi/stock && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc go build main.go
