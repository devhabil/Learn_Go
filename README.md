# Learn_Go

Repositori ini berisi kumpulan kode sederhana untuk mempelajari dan mempraktikkan konsep-konsep dasar dalam bahasa pemrograman Go.

## Struktur Proyek

Proyek ini berisi beberapa file `.go` di dalam paket `main` untuk mendemonstrasikan fungsionalitas yang berbeda.

- `import.go`: Menunjukkan cara mengimpor paket lokal (`helper`) dan menggunakan fungsi yang ada di dalamnya.
- `init.go`: Mendemonstrasikan penggunaan fungsi `init` yang dieksekusi secara otomatis saat sebuah paket diinisialisasi. File ini juga menunjukkan penggunaan _blank import_ (`_`) untuk mengimpor paket (`internal`) hanya untuk efek sampingnya (menjalankan fungsi `init`-nya).
- `number.go`: Contoh dasar untuk mencetak berbagai tipe data numerik di Go.
- `helper/`: Paket lokal yang berisi fungsi `SayHello`.
- `database/`: Paket lokal yang digunakan untuk mendemonstrasikan fungsi `init`.
- `internal/`: Paket yang dimaksudkan untuk penggunaan internal, juga untuk mendemonstrasikan _blank import_.

## Cara Menjalankan Contoh

Dalam Go, sebuah program yang dapat dieksekusi hanya boleh memiliki satu fungsi `main` di dalam paket `main`. Karena proyek ini memiliki beberapa file dengan fungsi `main` di paket yang sama, Anda tidak dapat menjalankannya secara bersamaan dengan `go run .`.

Anda harus menjalankan setiap file secara individual.

```bash
# Menjalankan contoh import
go run import.go

# Menjalankan contoh inisialisasi paket
go run init.go

# Menjalankan contoh angka
go run number.go
```

## Konsep yang Dipelajari

- Struktur dasar program Go (`package main`, `func main`).
- Membuat dan menggunakan paket kustom.
- Mengimpor paket (standar dan lokal).
- Fungsi `init` dan urutan inisialisasi paket.
- _Blank Import_ (`_`) untuk efek samping inisialisasi.
- Aturan mengenai satu fungsi `main` per program.
