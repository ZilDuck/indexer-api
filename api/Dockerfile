FROM golang:alpine AS builder

RUN adduser -D -u 1001 -g '' appuser
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go env CGO_ENABLED
RUN go mod download -x
RUN go mod verify
RUN go run ./cmd/internal/injectDependencies/main.go ./generated

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/api ./cmd/api

RUN chmod u+x /go/bin/*

FROM alpine:3.15.1

WORKDIR /app

RUN mkdir /app/logs

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/api /app/api