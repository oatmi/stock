# step-1: builder
FROM --platform=linux/x86_64 golang:1.20.1 as builder

RUN apt-get update \
    && apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross git

WORKDIR ./src/github.com/oatmi/stock
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc go build -o /app/stock

# step-2
FROM --platform=linux/x86_64 alpine:latest as runner
COPY --from=builder /app/stock /app/stock
EXPOSE 8888
ENTRYPOINT ["/app/stock"]
