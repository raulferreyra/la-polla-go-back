# 🐳 DEV ONLY
FROM golang:1.25-alpine

WORKDIR /app

# utilidades ligeras
RUN apk add --no-cache bash curl ca-certificates tzdata

# copiar archivos de módulos primero para cachear dependencias
COPY go.mod go.sum ./
RUN go mod download

# copiar el resto del código
COPY . .

# variables por defecto (se pueden overridear en docker-compose)
ENV APP_ENV=dev \
    HTTP_ADDR=:8080 \
    DATABASE_URL=postgres://postgres:postgres@db:5432/polla?sslmode=disable \
    JWT_SECRET=dev_jwt_secret_please_change \
    AES_KEY=0123456789abcdef0123456789abcdef

EXPOSE 8080

# correr en modo desarrollo (recompila al iniciar el contenedor)
CMD ["go", "run", "./cmd/api"]
