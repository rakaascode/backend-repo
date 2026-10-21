# Article Service Submission

REST API pengelolaan artikel menggunakan Go, Gin, GORM, dan MySQL.

## Prasyarat

- Go 1.26.0 atau versi yang kompatibel dengan `go.mod`.
- MySQL yang sudah berjalan dan akun yang dapat membuat database serta tabel.

Semua perintah berikut dijalankan dari direktori `article-service` yang berisi `go.mod`.

## 1. Siapkan konfigurasi

Salin contoh konfigurasi:

```bash
cp .env.example .env
```

Edit `.env` sesuai koneksi MySQL lokal:

```dotenv
DB_USER=root
DB_PASSWORD=ganti_password_mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=db_article
```

Isi `DB_PASSWORD` dengan password akun MySQL. Jika akun tidak menggunakan password, gunakan `DB_PASSWORD=`. File `.env` bersifat lokal; gunakan `.env.example` sebagai contoh konfigurasi.

## 2. Buat database dan tabel

Jalankan MySQL/MariaDB menggunakan XAMPP atau instalasi lainnya, lalu jalankan query berikut melalui terminal, phpMyAdmin, atau aplikasi database pilihan Anda:

```sql
CREATE DATABASE db_article;
USE db_article;

CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status VARCHAR(100) NOT NULL
);
```

Sesuaikan `.env` dengan akun MySQL yang digunakan dan gunakan `DB_NAME=db_article`. Jika akun tidak menggunakan password, isi `DB_PASSWORD=`.

## 3. Instal dependensi dan jalankan

```bash
go mod download
go run .
```

Server berjalan di **http://localhost:8081**. Setelah database berhasil terhubung, terminal menampilkan:

```text
Database connected and migrated successfully!
```

Biarkan terminal terbuka selama menggunakan API. Tekan `Ctrl+C` untuk menghentikan server.

Untuk menjalankan binary hasil build:

```bash
go build -o bin/article-service .
./bin/article-service
```

Jalankan binary dari direktori yang berisi `.env`, atau sediakan variabel lingkungan `DB_*` melalui shell.

## 4. Cek aplikasi

Dari terminal lain:

```bash
curl http://localhost:8081/health
curl http://localhost:8081/health/db
```

Kedua endpoint mengembalikan HTTP `200` dan status `UP` jika aplikasi serta koneksi database sehat.

## Endpoint artikel

| Method | URL | Fungsi |
| --- | --- | --- |
| POST | `/article/` | Membuat artikel |
| GET | `/article/10/0` | Mengambil maksimal 10 artikel dengan offset 0 |
| GET | `/article/1` | Mengambil artikel dengan ID 1 |
| PUT | `/article/1` | Memperbarui artikel dengan ID 1 |
| DELETE | `/article/1` | Menghapus artikel dengan ID 1 |

Contoh membuat artikel:

```bash
curl -X POST http://localhost:8081/article/ \
  -H 'Content-Type: application/json' \
  -d '{
    "title": "Contoh judul artikel pertama",
    "content": "Ini adalah contoh konten artikel untuk mencoba API secara lokal. Konten dibuat cukup panjang agar bisa dipakai sebagai contoh permintaan yang lengkap. Setelah berhasil dibuat, artikel dapat diambil, diperbarui, dan dihapus melalui endpoint yang tersedia.",
    "category": "Teknologi",
    "status": "Draft"
  }'
```

Respons create berbentuk `{"message":"Article created successfully","data":{...}}`. Gunakan ID dari `data.id` untuk request berikutnya.

## Kendala umum

- **Koneksi database gagal:** pastikan MySQL berjalan, database sudah dibuat, dan nilai `DB_*` sesuai. Akun harus memiliki izin mengakses database serta melakukan migrasi tabel.
- **Port 8081 sedang digunakan:** hentikan proses lain yang memakai port tersebut. Port server saat ini ditetapkan di `main.go`.
- **File .env tidak ditemukan:** jalankan aplikasi dari direktori `article-service` atau sediakan konfigurasi melalui variabel lingkungan.
- **Warning Gin debug mode:** ini normal saat pengembangan. Gunakan `GIN_MODE=release go run .` jika ingin menjalankan Gin dalam mode release.
