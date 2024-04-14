FROM golang:1.22-alpine

WORKDIR /app

COPY . /app

RUN go install github.com/cosmtrek/air@latest

RUN go mod download

CMD ["air", "-c", ".air.toml"]