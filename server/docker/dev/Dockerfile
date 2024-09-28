FROM golang:1.23.1-alpine
WORKDIR /app
COPY ../../ .

RUN go build -o app

ENTRYPOINT ./app