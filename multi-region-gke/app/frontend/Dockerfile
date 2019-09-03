FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /go/src/mypackage/myapp/
COPY . .
RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app
RUN chmod a+x /go/bin/app

FROM alpine
WORKDIR /go/bin
RUN mkdir -p /go/bin/templates
RUN mkdir -p /go/bin/assets/css
COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /go/src/mypackage/myapp/templates/* /go/bin/templates/
COPY --from=builder /go/src/mypackage/myapp/assets/css/* /go/bin/assets/css/
ENTRYPOINT ["/go/bin/app"]
