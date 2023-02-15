FROM golang:1.20.0-bullseye

WORKDIR /go/src/backend

COPY go.mod go.sum ./

RUN apt-get update \
    && apt-get install git \
    && go get github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql

CMD ["go", "run", "."]