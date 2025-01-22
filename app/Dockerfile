FROM golang:1.23.2-alpine AS builder

WORKDIR /usr/app

COPY ./go.mod ./go.sum ./
RUN go mod download
RUN go mod tidy

COPY ./entrypoint.sh /usr/app/entrypoint.sh
RUN chmod +x entrypoint.sh

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

ENTRYPOINT ["sh", "/usr/app/entrypoint.sh"]