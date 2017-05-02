FROM golang:alpine
RUN apk add --update git
RUN go get -u -v github.com/netroby/mock-server
EXPOSE 7890
ADD www /
CMD ["/go/bin/mock-server", "-w", "/www"]
