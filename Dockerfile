
FROM golang:1.20.3-alpine3.17 as builder
RUN mkdir -p /app
WORKDIR /app

COPY ./src/ /app
RUN go work sync
WORKDIR /app/cmd/lol
RUN go mod tidy
RUN go mod download 
WORKDIR /app/internal/config
RUN go mod tidy
RUN go mod download 
WORKDIR /app
RUN go build -ldflags "-s -w" -o /app/lol ./cmd/lol/main.go 

FROM node as nodejs
RUN mkdir /app
COPY ./ui/ /app/ui
WORKDIR /app/ui
RUN npm run build --production

FROM alpine:3.17.3
COPY --from=builder /app/lol /go/boostchickenlol
RUN mkdir /app/boostchickenlol
COPY --from=node /app /go/ui/
COPY ui /go
WORKDIR /go

LABEL org.opencontainers.image.maintainer="John Dorman <john@boostchicken.dev>"     
LABEL org.opencontainers.image.authors="John Dorman <john@boostchicken.dev>"        
LABEL org.opencontainers.image.title="boostchicken/lol"                          
LABEL org.opencontainers.image.vendor="boostchicken.dev"              
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="https://www.github.com/boostchicken/lol"       
LABEL org.opencontainers.image.source="https://www.github.com/boostchicken/lol"          
LABEL org.opencontainers.image.documentation="https://www.github.com/boostchicken/lol/blob/main/README.md"
LABEL org.opencontainers.image.description="bunnylol clone in go" 

COPY LICENSE /go/
COPY README.md /go/
ENTRYPOINT [ "/go/boostchickenlol" ]
CMD [ "bash"]