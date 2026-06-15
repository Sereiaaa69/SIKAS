# 🎓 SIKAS - Sistem Informasi Kas Mahasiswa

Aplikasi CLI berbasis **Go** untuk mengelola pembayaran kas mahasiswa per bulan.  
Dibuat untuk memenuhi Tugas Besar mata kuliah **Algoritma dan Pemrograman 2**.

---

## ✨ Fitur Aplikasi

| No | Fitur | Keterangan |
|----|-------|-------------|
| 1 | Tambah Mahasiswa | Input NIM dan Nama |
| 2 | Ubah Mahasiswa | Ubah nama berdasarkan NIM |
| 3 | Hapus Mahasiswa | Hapus data berdasarkan NIM |
| 4 | Catat Pembayaran | Input bulan & tanggal bayar |
| 5 | Sequential Search | Cari mahasiswa belum lunas (linear) |
| 6 | Binary Search | Cari mahasiswa belum lunas (bagi dua) |
| 7 | Insertion Sort | Urutkan mahasiswa berdasarkan nama (A-Z) |
| 8 | Selection Sort | Urutkan berdasarkan tunggakan (terbesar ke terkecil) |
| 9 | Statistik Kas | Total kas, jumlah lunas/belum lunas |
| 10 | Tampilkan Data | Lihat semua data mahasiswa |
| 0 | Keluar | Exit program |

---

## 🚀 Cara Menjalankan

### Prasyarat
- Install **Go** minimal versi 1.18  
  [Download Go](https://go.dev/dl/)

### Langkah-langkah

```bash
# Clone repository ini
git clone https://github.com/Sereiaaa69/SIKAS.git

# Masuk ke folder
cd SIKAS

# Jalankan program
go run main.go

```
## Struktur Program

```
SIKAS/
├── main.go          # Source code utama
└── README.md        # Dokumentasi

```
## Struktur Data

```go
type Mahasiswa struct {
    NIM          string
    Nama         string
    Pembayaran   [12]bool
    TanggalBayar [12]string
}
```
## Algoritma yang Digunakan

| Algoritma | Fungsi | Cara Kerja |
|-----------|--------|-------------|
| Sequential Search | Mencari mahasiswa belum lunas | Mengecek satu per satu dari awal sampai ketemu |
| Binary Search | Mencari mahasiswa belum lunas | Membagi data menjadi dua bagian, lalu mencari di bagian yang sesuai |
| Insertion Sort | Mengurutkan berdasarkan nama | Mengambil satu per satu data lalu menyisipkan ke posisi yang tepat |
| Selection Sort | Mengurutkan berdasarkan tunggakan | Mencari nilai tertinggi lalu menukarnya ke posisi awal |

## Kelompok

| Nama | NIM |
|------|-----|
| Fathan Ibrahamofiq | 108072500037 |
| Muhammad Taufiqa Hadi | 108072500163 |

**Kelas:** IF-05-02

## Dosen Pengampu

- Kharisma Monika Dian Pertiwi, S.Kom., M.Kom.
- Dr. Tanzilal Mustaqim, S.Kom., M.Kom.

## Universitas

**Program Studi S1 Informatika**  
**Fakultas Informatika**  
**Universitas Telkom**  
**Tahun Akademik 2025/2026**

