FROM golang:1.23.5

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
COPY ./omikuji .

RUN go build -o omikuji

CMD ["./omikuji"]
