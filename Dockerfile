FROM golang:1.19-alpine as builder
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /tmp

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go mod download && go build -a -o /go/bin/boostchickenlol


RUN 

FROM alpine:latest
EXPOSE 80
EXPOSE 8080
VOLUME /opt/boostchickenlol/config.yaml
COPY --from=builder /go/bin/boostchickenlol /opt/boostchickenlol/boostchickenlol
COPY config.yaml /opt/boostchickenlol/config.yaml

ENTRYPOINT [ "/opt/boostchickenlol/boostchickenlol" ]