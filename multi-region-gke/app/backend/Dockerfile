FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /go/src/mypackage/myapp/
COPY . .
RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app
RUN chmod a+x /go/bin/app

FROM alpine
WORKDIR /go/bin
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
