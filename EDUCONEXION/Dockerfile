# 1. Etapa de build
FROM golang:1.24-alpine AS build

WORKDIR /app

# Copiamos módulos y descargamos dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiamos TODO el código de la aplicación
COPY . .

# Compilamos el binario
RUN go build -o server ./cmd/main.go

# 2. Etapa de ejecución
FROM alpine:latest

WORKDIR /app

# Traemos únicamente el binario compilado
COPY --from=build /app/server .

# Exponemos el puerto en el que corre Gin
EXPOSE 3000

# Comando por defecto
CMD ["./server"]
