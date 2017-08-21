FROM golang:1.6

RUN mkdir /helloworld
WORKDIR /helloworld
ADD . /helloworld

ENV GOPATH /helloworld

RUN go get github.com/mitchellh/gox

RUN $GOPATH/bin/gox -osarch='linux/amd64'
CMD ./helloworld_linux_amd64
