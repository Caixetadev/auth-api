FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/auth-api ./cmd/api/

EXPOSE 3333

CMD ["/app/auth-api"]
