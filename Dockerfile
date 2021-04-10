FROM golang:alpine

WORKDIR /app

COPY . .

RUN go install
RUN go build -o ./server main.go

EXPOSE 8080

CMD ["./server"]