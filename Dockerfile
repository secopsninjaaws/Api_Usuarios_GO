# Etapa 1: build da aplicação
FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila o binário com base no main.go dentro da pasta /api
RUN go build -o api ./api/main.go

# Etapa 2: imagem final (mínima)
FROM alpine:latest

WORKDIR /app

# Copia apenas o binário gerado
COPY --from=builder /app/api .

# Expõe a porta que sua aplicação utiliza (ajuste se necessário)
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./api"]