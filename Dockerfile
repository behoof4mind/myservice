############################
# STEP 1 build executable binary
############################
FROM golang:1.15-alpine3.12 AS builder
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git=2.26.2-r0 ca-certificates=20191127-r4

WORKDIR $GOPATH/src/myservice/
COPY . .
# Fetch dependencies using go get
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /root/myservice/myservice
COPY ./models/migrations /root/myservice/models/migrations
COPY ./templates /root/myservice/templates
COPY ./public /root/myservice/public

############################
# STEP 2 build a small image
############################
FROM scratch
LABEL maintainer="dlavrushko@protonmail.com"
WORKDIR /app/
COPY --from=builder /root/myservice/ .

ENTRYPOINT ["/app/myservice"]