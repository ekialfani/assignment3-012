# Assignment 3

Proyek ini merupakan salah satu tugas MSIB Kampus merdeka di hacktiv8.

## Instalasi Manual

Berikut adalah langkah-langkah untuk menginstal dan menjalankan proyek ini secara manual.

### Persyaratan

Pastikan Anda telah menginstal Go (Golang) dan memiliki Go environment yang terkonfigurasi dengan baik di sistem Anda.

### Langkah 1: Clone Repository

Clone repository proyek ini ke dalam direktori lokal Anda:

```bash
git clone https://github.com/ekialfani/assignment3-012.git
cd assignment3-012
```

### Langkah 2: Instalasi Dependensi
```bash
go get -u gorm.io/gorm gorm.io/driver/postgres github.com/gin-gonic/gin github.com/joho/godotenv
```

### Langkah 3: Menjalankan Service Pertama
- Pindah ke direktori service pertama:
  ```bash
  cd main-service
  ```
- Buat file .env
  ```bash
  touch .env
  ```

- Tambahkan environment variable
  ```bash
    DB_HOST=Host database (contoh: localhost)
    DB_USER=Nama pengguna database (contoh: postgres)
    DB_PASSWORD=Kata sandi pengguna database (contoh: postgres)
    DB_PORT=Port database (contoh: 5432)
    DB_NAME=Nama database yang akan digunakan (contoh: "nama database")
    DB_SSLMODE=Mode SSL database (contoh: disable)
  ```
- Jalankan service pertama:
  ```bash
  go run main.go
  ```

### Langkah 4: Menjalankan Service Kedua untuk memperbarui data cuaca
Pindah ke direktori service kedua:
```bash
cd req-api-service
```

Jalankan service kedua:
```bash
go run main.go
```