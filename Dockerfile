FROM golang:alpine
RUN apk add --update git
RUN go get -u -v github.com/netroby/mock-server
EXPOSE 7890
CMD ["/go/bin/mock-server"]
