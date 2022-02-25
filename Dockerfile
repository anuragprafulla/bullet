FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git 

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ ./...

FROM scratch

COPY --from=builder /app/bin/cmd .

CMD ["./cmd"]