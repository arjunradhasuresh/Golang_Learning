FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./


RUN go mod download 


COPY . .


RUN go build -o main . 

EXPOSE 3000

CMD ["/go/bin/air", "-c", ".air.toml"]


