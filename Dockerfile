FROM golang:1.20.1-alpine3.17 as builder

RUN mkdir -p /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go mod download && go build -o /app/boostchickenlol

FROM alpine:3.17.2
COPY --from=builder /app/boostchickenlol /go/boostchickenlol
WORKDIR /go

CMD [ "/go/boostchickenlol" ]