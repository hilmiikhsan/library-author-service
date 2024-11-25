FROM golang:1.22.8-alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o library-author-service

RUN chmod +x library-author-service

EXPOSE 9092

EXPOSE 6002

CMD ["./library-author-service"]
