FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /go/src/outlet_client

COPY . .

RUN go mod init
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/outlet_client


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD order.json /app/order.json
COPY --from=builder /app/outlet_client .

CMD ["./outlet_client"]
