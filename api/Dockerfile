FROM golang:alpine AS builder

RUN apk add curl unzip

RUN curl -Ls https://releases.hashicorp.com/consul-template/0.29.0/consul-template_0.29.0_linux_amd64.zip -o /usr/local/bin/consul-template.zip
RUN cd /usr/local/bin && unzip /usr/local/bin/consul-template.zip
RUN chmod +x /usr/local/bin/consul-template

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

COPY --from=builder /usr/local/bin/consul-template /usr/local/bin/consul-template
COPY --from=builder /etc/passwd                    /etc/passwd
COPY --from=builder /go/bin/api                    /app/api
COPY                ./bin/run.sh                   /app/run.sh
