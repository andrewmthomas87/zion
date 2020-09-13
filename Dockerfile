FROM golang:1.15

WORKDIR /go/src/github.com/andrewmthomas87/zion
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["zion", "serve"]
