FROM golang:1.23.5

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
COPY ./hello-web .

RUN go build -o hello-web

EXPOSE 8080

CMD ["./hello-web"]
