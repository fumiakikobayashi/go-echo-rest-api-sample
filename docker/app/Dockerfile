FROM golang:1.21.4-bullseye

WORKDIR /app

COPY go.mod go.sum ./

RUN apt-get update \
    && apt-get install git \
    && go get github.com/labstack/echo/v4 \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/rs/zerolog/log \
    && go get github.com/go-delve/delve/cmd/dlv@latest   \