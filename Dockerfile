FROM golang:1.15-alpine AS dev

WORKDIR /bot

RUN apk add git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN go install github.com/ReeceDonovan/uni-bot

CMD [ "go", "run", "*.go" ]

FROM alpine

WORKDIR /bin

COPY --from=dev /go/bin/uni-bot ./uni-bot


CMD ["sh", "-c", "uni-bot -p"]
