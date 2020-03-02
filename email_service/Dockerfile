FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /go/src/email_srv

COPY . .

RUN go mod init
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/email_srv


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/email_srv .

CMD ["./email_srv"]
