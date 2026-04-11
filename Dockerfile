FROM golang:1.25.1-alpine3.20

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download

RUN addgroup -g 1001 appgroup && adduser -D -u 1001 -G appgroup appuser
USER appuser