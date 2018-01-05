FROM golang:latest
MAINTAINER wzy rpc test
ADD ./server.go /go/src
WORKDIR /go/src
RUN go build -o server server.go
RUN chmod +x server
EXPOSE 10012
CMD ["./server"]
