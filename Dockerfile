FROM golang:1.20.1-alpine3.17 as builder
ENV GOARCH=amd64
ENV GOOS=linux
RUN mkdir -p /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY config /app/config
RUN go mod download 
RUN go build 

FROM alpine:3.17.2
COPY --from=builder /app/lol /go/boostchickenlol
WORKDIR /go

CMD [ "/go/boostchickenlol" ]