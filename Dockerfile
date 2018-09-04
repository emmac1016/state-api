FROM golang:1.10
LABEL author="Emma Chirapongse <emmac1016@gmail.com>"

ADD . /go/src/github.com/emmac1016/state-api
WORKDIR /go/src/github.com/emmac1016/state-api

CMD ["./main"]
EXPOSE 8080