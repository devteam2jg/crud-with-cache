FROM golang:1.24

RUN mkdir -p /go/bin /go/logs /go/shared /app

WORKDIR /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go install github.com/cosmtrek/air@v1.49

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080