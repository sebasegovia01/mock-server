# Etapa de compilación
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copiar los archivos de dependencias primero
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa final
FROM alpine:latest

WORKDIR /app

# Copiar el binario compilado desde la etapa de builder
COPY --from=builder /app/main .

# Puerto que expone la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]