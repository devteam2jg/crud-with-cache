FROM golang:1.24

RUN mkdir -p /go/bin /go/logs /go/shared /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080
