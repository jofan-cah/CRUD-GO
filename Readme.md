# CRUD Karyawan dengan Go dan MySQL

Proyek ini adalah aplikasi CRUD sederhana untuk mengelola data karyawan menggunakan bahasa pemrograman Go dan database MySQL. Aplikasi ini mendukung fitur untuk menampilkan, menambah, memperbarui, dan menghapus data karyawan melalui antarmuka berbasis web.

## Fitur Utama
- **Menampilkan daftar karyawan**: Lihat semua data karyawan yang tersedia.
- **Menambah karyawan baru**: Tambahkan data karyawan melalui formulir.
- **Memperbarui data karyawan**: Edit data karyawan yang ada.
- **Menghapus karyawan**: Hapus data karyawan berdasarkan ID.

## Struktur Direktori
```
├── controller/
│   ├── karyawanController.go  # Logika bisnis untuk operasi CRUD karyawan
├── views/
│   ├── index.html             # Halaman untuk menampilkan daftar karyawan
│   ├── create.html            # Formulir untuk menambah karyawan
│   ├── update.html            # Formulir untuk memperbarui karyawan
├── main.go                    # Entry point aplikasi
├── routes.go                  # Definisi rute aplikasi
└── README.md                  # Dokumentasi proyek
```

## Prasyarat
Sebelum menjalankan proyek ini, pastikan Anda memiliki:
- Go (minimal versi 1.18)
- MySQL Server
- Paket Go berikut:
  - `github.com/go-sql-driver/mysql`
  - `github.com/gorilla/mux` (opsional jika ingin mendukung rute dinamis)

## Instalasi dan Pengaturan

1. **Clone repositori**:
   ```bash
   git clone https://github.com/username/karyawan-crud.git
   cd karyawan-crud
   ```

2. **Konfigurasi database**:
   Buat database MySQL dan tabel `karyawan`:
   ```sql
   CREATE DATABASE karyawan_db;

   USE karyawan_db;

   CREATE TABLE karyawan (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       npwp VARCHAR(20) NOT NULL,
       adddres TEXT NOT NULL
   );
   ```

3. **Konfigurasi koneksi database**:
   Edit file `main.go` dan sesuaikan string koneksi:
   ```go
   dsn := "user:password@tcp(127.0.0.1:3306)/karyawan_db"
   ```

4. **Jalankan aplikasi**:
   ```bash
   go run main.go
   ```

5. **Akses aplikasi**:
   Buka browser Anda dan akses [http://localhost:8080/karyawan](http://localhost:8080/karyawan).

## API Endpoint

| Method | Endpoint              | Deskripsi                     |
|--------|-----------------------|-------------------------------|
| GET    | `/karyawan`           | Menampilkan daftar karyawan   |
| POST   | `/karyawan/create`    | Menambahkan karyawan baru     |
| POST   | `/karyawan/update`    | Memperbarui data karyawan     |
| POST   | `/karyawan/delete`    | Menghapus karyawan berdasarkan ID |

## Catatan Penting
Jika Anda menggunakan `http.ServeMux`, pastikan rute untuk delete menggunakan query string (`/karyawan/delete?id=<id>`). Jika ingin mendukung rute dinamis seperti `/karyawan/delete/{id}`, gunakan library seperti **Gorilla Mux**.

## Lisensi
Proyek ini menggunakan lisensi MIT. Anda bebas menggunakannya sesuai ketentuan lisensi.
