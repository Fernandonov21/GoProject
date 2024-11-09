# Usa una imagen base de Go
FROM golang:1.20-alpine AS build

# Establece el directorio de trabajo
WORKDIR /app

# Copia el código fuente al contenedor
COPY . .

# Instala dependencias y compila la aplicación
RUN go build -o app .

# Usa una imagen base mínima para ejecutar la aplicación
FROM alpine:latest

# Instala dependencias mínimas necesarias para Go (si las necesita)
RUN apk --no-cache add ca-certificates

# Copia el archivo binario desde la imagen de construcción
COPY --from=build /app/app /app/app

# Expone el puerto en el que la aplicación estará escuchando
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/app/app"]
