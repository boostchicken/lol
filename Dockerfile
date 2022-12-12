FROM golang:1.19-alpine
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir -p /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go mod download 
RUN go build -o /app/boostchickenlol

EXPOSE 8080
VOLUME /app/config.yaml
CMD [ "/app/boostchickenlol" ]