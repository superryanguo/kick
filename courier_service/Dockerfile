FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /go/src/courier_server
COPY . .

RUN go mod init
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/courier_server

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/courier_server .

CMD ["./courier_server"]
