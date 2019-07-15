FROM golang:1.12.1-stretch
WORKDIR $GOPATH/src/github.com/tkachenkom/tm-stream/
COPY . .
RUN go build -v ./cmd/main.go
EXPOSE 8081
CMD ["./main"]