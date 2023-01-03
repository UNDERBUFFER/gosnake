FROM golang:1.19 AS build

WORKDIR /app

COPY ./src/* ./
RUN go mod download
