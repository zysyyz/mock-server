FROM golang:alpine
RUN go get -u -v github.com/netroby/mock-server

CMD ["/go/bin/mock-server"]
