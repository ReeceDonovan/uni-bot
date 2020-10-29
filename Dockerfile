FROM golang:1.15-alpine AS dev

WORKDIR /bot

RUN apk add git

RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN go install github.com/ReeceDonovan/CS-bot

CMD [ "go", "run", "*.go" ]

FROM alpine

WORKDIR /bin

COPY --from=dev /go/bin/CS-bot ./CS-bot

EXPOSE 8888:80


CMD ["sh", "-c", "CS-bot -p"]
