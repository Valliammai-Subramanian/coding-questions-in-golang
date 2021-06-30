FROM golang:1.14.3-alpine AS build

COPY . ./src

RUN go build /src/.
