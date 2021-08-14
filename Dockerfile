FROM golang:1.16-alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/app/run .

CMD ["./run"]