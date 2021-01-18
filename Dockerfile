FROM golang:1.14-alpine

WORKDIR /go/src/demo-server
COPY . .

RUN go build main.go

CMD ["./main"]