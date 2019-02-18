FROM golang:1.11.4
ENV GO111MODULE=on

WORKDIR /go/src/github.com/gmohre/xkcd
COPY . .
RUN go mod download

RUN go build
CMD go get github.com/oxequa/realize && realize start

EXPOSE 8000


