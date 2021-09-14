FROM golang:alpine

WORKDIR /go/src

COPY . .

CMD ["top"]