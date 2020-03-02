# Builder
FROM golang:1.12.8-alpine3.10 as builder

RUN apk update && apk upgrade && \
    apk --no-cache --update add git make && \
    go get -u github.com/golang/dep/cmd/dep


WORKDIR /app

COPY . .

RUN make engine
#dep init -v && build -o engine app/main.go

## Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8084

COPY --from=builder /app/engine /app

CMD /app/engine