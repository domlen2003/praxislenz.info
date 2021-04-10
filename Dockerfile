FROM golang:alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go clean --modcache

RUN go build -o ./server main.go

EXPOSE 8080

CMD ["./server"]
