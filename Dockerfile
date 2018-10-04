
# FROM golang:1.9
# COPY . /go/src/github.com/davitron/go-micro-services
# WORKDIR /go/src/github.com/davitron/go-micro-services
# RUN go install -ldflags="-s -w" ./cmd/...

FROM golang:1.9

WORKDIR /go/src/github.com/davitron/go-micro-services

COPY . /go/src/github.com/davitron/go-micro-services

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN go get -u google.golang.org/grpc

RUN go get -u github.com/golang/protobuf/protoc-gen-go

# RUN for f in services/**/proto/*.proto; do protoc --go_out=plugins=grpc:. $$f; done

RUN go install -ldflags="-s -w" ./cmd/...

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# RUN mkdir /app
# WORKDIR /app

# COPY --from=builder /go/src/github.com/davitron/go-micro-services .