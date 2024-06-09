FROM golang:1.22.4-alpine

WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o main .

CMD ["./main"]
