FROM golang:1.14

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# RUN go get -d -v ./...
# RUN go install -v ./...
# RUN go mod tidy
# RUN go build

ENTRYPOINT ["docker/entrypoint.sh"]