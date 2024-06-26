
FROM golang:1.22.2-alpine3.18 as builder
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

FROM node:22-slim AS nodejs
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable && corepack prepare pnpm@latest --activate
RUN mkdir /app
COPY ./ui/ /app/ui
WORKDIR /app/ui
RUN pnpm build


FROM alpine:3
RUN mkdir /go
COPY --from=builder /app/lol /go/boostchickenlol
COPY --from=nodejs /app/ui/ /go/ui/
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

ENTRYPOINT [ "/go/boostchickenlol" ]
CMD [ "bash"]
