FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

WORKDIR /app/cmd

RUN GOOS=linux go build .

RUN chmod +x ./cmd

ENTRYPOINT ["./cmd"]