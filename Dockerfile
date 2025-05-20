FROM golang:1.23 as builder

WORKDIR /app

# Копируем сначала файлы модулей
COPY backend/go.mod backend/go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем остальные файлы бекенда
COPY backend/ ./

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o hospital-management ./cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/hospital-management .
# Теперь копируем фронтенд из корня проекта
COPY frontend ./frontend

# УДАЛИТЬ все ENV переменные из Dockerfile (оставить только PORT)
ENV PORT=8080

EXPOSE 8080

CMD ["./hospital-management"]