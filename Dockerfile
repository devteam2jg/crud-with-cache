FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /go/bin /go/logs /go/shared && \
    go build -o /go/bin/app

EXPOSE 8080

CMD ["/go/bin/app"]