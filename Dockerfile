FROM golang:1.23.4-alpine AS builder

RUN apk update && apk add --no-cache gcc musl-dev

# Atur direktori kerja di dalam container.
WORKDIR /app

# Salin go.mod dan go.sum terlebih dahulu untuk memanfaatkan caching Docker.
# Ini agar dependensi tidak diunduh ulang jika hanya kode sumber yang berubah.
COPY go.mod go.sum ./

# Unduh semua dependensi Golang.
RUN go mod download && go mod tidy

# Salin seluruh kode sumber proyek ke dalam direktori kerja.
COPY . .

# Kompilasi aplikasi Golang menjadi binary executable.
# -o server: Nama output binary executable menjadi 'server'.
# cmd/server/main.go: Path ke file main aplikasi.
RUN go build -o server cmd/server/main.go

# PENTING: Instal CLI tool 'migrate' agar bisa digunakan untuk migrasi database.
# Binary ini akan diinstal di GOBIN default, biasanya /go/bin.
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# --- Tahap 2: Final Image (The 'runtime' Stage) ---
# Menggunakan image Alpine Linux yang sangat ringan sebagai base image akhir.
FROM alpine:latest

# Atur direktori kerja di dalam container akhir.
WORKDIR /app

# Tambahkan data zona waktu agar fungsi waktu di Go bekerja dengan benar.
# Kemudian set zona waktu ke Asia/Jakarta.
RUN apk update && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime \
    && echo "Asia/Jakarta" > /etc/timezone \
    && rm -rf /var/cache/apk/* # Bersihkan cache apk setelah instalasi

# Salin binary executable 'server' yang sudah dikompilasi dari tahap 'builder'.
# Hanya binary yang disalin, tidak ada source code atau tool build.
COPY --from=builder /app/server .

COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

COPY --from=builder /app/credentials ./credentials

COPY --from=builder /app/database/migration ./database/migration

CMD ["./server"]
