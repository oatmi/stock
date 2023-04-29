# builder
FROM alpine AS builder
RUN apk add --no-cache --update go

WORKDIR ./src/github.com/oatmi/stock
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/stock
COPY ./template/* /app/template/
COPY ./sdk/* /app/sdk/

# runner
FROM alpine as runner
COPY --from=builder /app/stock /app/stock
COPY --from=builder /app/template/* /app/template/
COPY --from=builder /app/sdk/* /app/sdk/
WORKDIR /app
EXPOSE 8888
ENTRYPOINT ["./stock"]
