FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 5000

CMD ["./main"]
