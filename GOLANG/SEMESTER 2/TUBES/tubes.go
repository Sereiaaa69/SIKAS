package main

import (
	"fmt"
	"strings"
)

const MAX = 100

type Mahasiswa struct {
	NIM          string
	Nama         string
	Pembayaran   [12]bool
	TanggalBayar [12]string
}

var data [MAX]Mahasiswa
var n int
var kasBulanan int

// ================= HITUNG =================

func hitungBulanBayar(m Mahasiswa) int {
	jumlah := 0
	for i := 0; i < 12; i++ {
		if m.Pembayaran[i] {
			jumlah++
		}
	}
	return jumlah
}

func hitungTunggakan(m Mahasiswa) int {
	return (12 - hitungBulanBayar(m)) * kasBulanan
}

// ================= CRUD =================

func tambahMahasiswa() {
	if n >= MAX {
		fmt.Println("Data penuh!")
		return
	}

	fmt.Print("NIM : ")
	fmt.Scan(&data[n].NIM)

	fmt.Print("Nama : ")
	fmt.Scan(&data[n].Nama)

	n++
	fmt.Println("Data berhasil ditambahkan")
}

func ubahMahasiswa() {
	var nim string

	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)

	for i := 0; i < n; i++ {
		if data[i].NIM == nim {

			fmt.Print("Nama baru : ")
			fmt.Scan(&data[i].Nama)

			fmt.Println("Data berhasil diubah")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func hapusMahasiswa() {
	var nim string

	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)

	for i := 0; i < n; i++ {
		if data[i].NIM == nim {

			for j := i; j < n-1; j++ {
				data[j] = data[j+1]
			}

			n--
			fmt.Println("Data berhasil dihapus")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

// ================= PEMBAYARAN =================

func catatPembayaran() {
	var nim string
	var bulan int
	var tanggal string

	fmt.Print("NIM : ")
	fmt.Scan(&nim)

	for i := 0; i < n; i++ {
		if data[i].NIM == nim {

			fmt.Print("Bulan (1-12) : ")
			fmt.Scan(&bulan)

			if bulan < 1 || bulan > 12 {
				fmt.Println("Bulan tidak valid")
				return
			}

			fmt.Print("Tanggal Bayar : ")
			fmt.Scan(&tanggal)

			data[i].Pembayaran[bulan-1] = true
			data[i].TanggalBayar[bulan-1] = tanggal

			fmt.Println("Pembayaran berhasil dicatat")
			return
		}
	}

	fmt.Println("Mahasiswa tidak ditemukan")
}

// ================= TAMPIL =================

func tampilkanData() {

	fmt.Println("\n=== DATA MAHASISWA ===")

	for i := 0; i < n; i++ {

		fmt.Println("-------------------")
		fmt.Println("NIM :", data[i].NIM)
		fmt.Println("Nama :", data[i].Nama)

		bayar := hitungBulanBayar(data[i])
		fmt.Println("Sudah Bayar :", bayar, "bulan")
		fmt.Println("Tunggakan :", hitungTunggakan(data[i]))

		if bayar == 12 {
			fmt.Println("Status : LUNAS")
		} else {
			fmt.Println("Status : BELUM LUNAS")
		}
	}
}

// ================= SEARCH =================

// Sequential Search (BELUM LUNAS)
func sequentialSearchBelumLunas() {

	fmt.Println("\n=== SEQUENTIAL SEARCH BELUM LUNAS ===")

	ketemu := false

	for i := 0; i < n; i++ {
		if hitungBulanBayar(data[i]) < 12 {
			fmt.Println(data[i].NIM, "-", data[i].Nama)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Semua mahasiswa sudah lunas")
	}
}

// Binary Search (BELUM LUNAS)
func binarySearchBelumLunas() {

	// STEP 1: pindahkan BELUM LUNAS ke depan (sorting sederhana)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {

			if hitungBulanBayar(data[i]) == 12 && hitungBulanBayar(data[j]) < 12 {
				data[i], data[j] = data[j], data[i]
			}
		}
	}

	// STEP 2: binary search di area BELUM LUNAS
	low := 0
	high := n - 1

	for low <= high {

		mid := (low + high) / 2

		if hitungBulanBayar(data[mid]) < 12 {

			fmt.Println("\n=== DITEMUKAN BELUM LUNAS ===")
			fmt.Println("NIM :", data[mid].NIM)
			fmt.Println("Nama :", data[mid].Nama)
			fmt.Println("Tunggakan :", hitungTunggakan(data[mid]))
			return
		}

		low = mid + 1
	}

	fmt.Println("Tidak ada mahasiswa belum lunas")
}

// ================= SORT =================

func insertionSortNama() {

	for i := 1; i < n; i++ {

		key := data[i]
		j := i - 1

		for j >= 0 && strings.ToLower(data[j].Nama) > strings.ToLower(key.Nama) {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}

	fmt.Println("\n=== HASIL SORT NAMA ===")

	for i := 0; i < n; i++ {
		fmt.Println(data[i].NIM, "-", data[i].Nama)
	}
}

func selectionSortTunggakan() {

	for i := 0; i < n-1; i++ {

		max := i

		for j := i + 1; j < n; j++ {
			if hitungTunggakan(data[j]) > hitungTunggakan(data[max]) {
				max = j
			}
		}

		data[i], data[max] = data[max], data[i]
	}

	fmt.Println("\n=== HASIL SORT TUNGGAKAN ===")

	for i := 0; i < n; i++ {
		fmt.Println(data[i].NIM, "-", data[i].Nama, "-", hitungTunggakan(data[i]))
	}
}

// ================= STATISTIK =================

func statistikKas() {

	total := 0
	lunas := 0

	for i := 0; i < n; i++ {

		bayar := hitungBulanBayar(data[i])
		total += bayar * kasBulanan

		if bayar == 12 {
			lunas++
		}
	}

	fmt.Println("\n=== STATISTIK ===")
	fmt.Println("Total Kas :", total)
	fmt.Println("Mahasiswa :", n)
	fmt.Println("Lunas :", lunas)
	fmt.Println("Belum Lunas :", n-lunas)
}

// ================= MAIN =================

func main() {

	fmt.Println("=== SISTEM KAS MAHASISWA ===")

	fmt.Print("Kas per bulan : ")
	fmt.Scan(&kasBulanan)

	pilihan := -1

	for pilihan != 0 {

		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tambah")
		fmt.Println("2. Ubah")
		fmt.Println("3. Hapus")
		fmt.Println("4. Pembayaran")
		fmt.Println("5. Sequential Search Belum Lunas")
		fmt.Println("6. Binary Search Belum Lunas")
		fmt.Println("7. Insertion Sort Nama")
		fmt.Println("8. Selection Sort Tunggakan")
		fmt.Println("9. Statistik")
		fmt.Println("10. Tampilkan Data")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih : ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahMahasiswa()
		case 2:
			ubahMahasiswa()
		case 3:
			hapusMahasiswa()
		case 4:
			catatPembayaran()
		case 5:
			sequentialSearchBelumLunas()
		case 6:
			binarySearchBelumLunas()
		case 7:
			insertionSortNama()
		case 8:
			selectionSortTunggakan()
		case 9:
			statistikKas()
		case 10:
			tampilkanData()
		case 0:
			fmt.Println("Keluar program")
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}