FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/http/main.go

EXPOSE 8080

CMD ["/app/main"]
