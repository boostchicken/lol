FROM golang:1.19-alpine
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go mod download 
RUN go build -o /boostchickenlol

EXPOSE 8080
CMD [ "/boostchickenlol" ]