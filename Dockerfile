# Use an official Golang runtime as a base image
FROM golang:1.22.0-alpine3.18 AS builder

# Set the working directory in the container
WORKDIR /app

COPY go.mod go.sum ./

# Descarga e instala dependencias
RUN go mod download

# Copy the local package files to the container's workspace
COPY . .

#  Descarga e instala dependencias de terceros
# RUN go get -u github.com/gin-gonic/gin
# RUN go get -u gorm.io/gorm
# RUN go get -u gorm.io/driver/mysql
RUN CGO=ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Construye la APP de Go
# RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/.env .

CMD ["./main"]
