FROM golang:latest

WORKDIR /app
COPY . .
RUN go build -o go-api
CMD ["./go-api"]
