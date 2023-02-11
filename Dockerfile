FROM golang:1.20-alpine as builder
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir -p /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go mod download && go build -o /app/boostchickenlol

FROM golang:1.20-alpine
COPY --from=builder /app/boostchickenlol /go/boostchickenlol
COPY config.yaml /go

CMD [ "/go/boostchickenlol" ]