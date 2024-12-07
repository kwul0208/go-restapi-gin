# image go untuk build aplikasi
FROM golang:1.23.0 AS builder

# set work dir didalam container
WORKDIR /app

# copy semua file aplikasi
COPY . .

# Download dependecies dan build aplikasi
RUN go mod tidy
RUN go build -o app


# set port default untuk aplikasi
EXPOSE 8080

# perintah untuk menjalankan aplikasi
CMD ["./app"]
