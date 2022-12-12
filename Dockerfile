FROM golang:1.19-alpine as builder


COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/boostchickenlol

FROM alpine:latest
EXPOSE 80
EXPOSE 8080
VOLUME /opt/boostchickenlol/config.yaml
COPY --from=builder /go/bin/boostchickenlol /opt/boostchickenlol/boostchickenlol -
COPY config.yaml /opt/boostchickenlol/config.yaml

ENTRYPOINT [ "/opt/boostchickenlol/boostchickenlol" ]