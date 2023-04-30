FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/crnt-auth-service ./cmd/crnt-auth-service/main.go

CMD ["./bin/crnt-auth-service"]
