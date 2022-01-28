FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git gcc g++ make libc-dev pkgconfig zeromq-dev curl libunwind-dev

RUN mkdir -p $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

RUN go env CGO_ENABLED
RUN go mod download -x
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/api ./cmd/api

RUN chmod u+x /go/bin/*

FROM scratch

COPY --from=builder /go/bin/api /app/api

WORKDIR /app

ENTRYPOINT ["/app/api"]
