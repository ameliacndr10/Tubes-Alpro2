package main

import (
	"fmt"
	"strings"
)

func login() bool {
	const passwordBenar = "admin123"
	var password string
	fmt.Println("=== Login Aplikasi Pengelolaan Sampah ===")
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	if password != passwordBenar {
		fmt.Println("Password salah. AKses ditolak.")
		return false
	}
	return true
}

type Sampah struct {
	ID              int
	Jenis           string
	Jumlah          int
	MetodeDaurUlang string
}

var dataSampah []Sampah

func main() {
	if !login() {
		return
	}

	dataSampah = []Sampah{
		{ID: 1, Jenis: "Plastik", Jumlah: 10, MetodeDaurUlang: "Recycle(daurulang)"},
		{ID: 2, Jenis: "Kertas", Jumlah: 5, MetodeDaurUlang: "Reduce(mengurangi)"},
		{ID: 3, Jenis: "Logam", Jumlah: 8, MetodeDaurUlang: "Recycling(dilebur)"},
	}

	var pilihan int

	for {
		fmt.Println("\n=== Aplikasi Pengelolaan Data Sampah ===")
		fmt.Println("1. Tambah Data Sampah")
		fmt.Println("2. Ubah Data Sampah")
		fmt.Println("3. Hapus Data Sampah")
		fmt.Println("4. Cari Data Sampah")
		fmt.Println("5. Urutkan Data Sampah")
		fmt.Println("6. Tampilkan Statistik")
		fmt.Println("7. Tampilkan Semua Data")
		fmt.Println("8. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData()
		case 2:
			ubahData()
		case 3:
			hapusData()
		case 4:
			cariData()
		case 5:
			urutkanData()
		case 6:
			tampilkanStatistik()
		case 7:
			tampilkan()
		case 8:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
func idSama(id int) bool {

	for _, s := range dataSampah {
		if s.ID == id {
			return true
		}
	}
	return false
}
func tambahData() {

	for {
		var jenis, metode string
		var id, jumlah int

		fmt.Print("Masukkan ID sampah: ")
		fmt.Scan(&id)

		if idSama(id) {
			fmt.Println("ID sudah ada, silahkan masukkan ID lain.")
			continue
		}

		fmt.Print("Masukkan jenis sampah: ")
		fmt.Scan(&jenis)
		fmt.Print("Masukkan jumlah sampah: ")
		fmt.Scan(&jumlah)
		fmt.Print("Masukkan metode daur ulang: ")
		fmt.Scan(&metode)

		dataSampah = append(dataSampah, Sampah{
			ID:              id,
			Jenis:           jenis,
			Jumlah:          jumlah,
			MetodeDaurUlang: metode})
		if konfirmasi() {
			break
		}
		fmt.Println("Data berhasil ditambahkan")
	}

}

func ubahData() {

	for {
		if len(dataSampah) == 0 {
			fmt.Println("Belum ada data untuk diubah.")
			return
		}

		tampilkanData()

		var index int
		fmt.Print("Masukkan nomor data yang ingin diubah: ")
		fmt.Scan(&index)

		if index > 0 && index <= len(dataSampah) {
			var jenis, metode string
			var jumlah int

			fmt.Print("Masukkan jenis sampah baru: ")
			fmt.Scan(&jenis)
			fmt.Print("Masukkan jumlah sampah baru: ")
			fmt.Scan(&jumlah)
			fmt.Print("Masukkan metode daur ulang: ")
			fmt.Scan(&metode)

			dataSampah[index-1] = Sampah{Jenis: jenis, Jumlah: jumlah, MetodeDaurUlang: metode}
			fmt.Println("Data berhasil diubah.")
		} else {
			fmt.Println("Nomor tidak valid.")
		}

		if konfirmasi() {
			break
		}
	}

}

func hapusData() {
	for {
		tampilkanData()
		var ID int
		fmt.Print("Masukkan nomor ID yang ingin dihapus: ")
		fmt.Scan(&ID)

		var idx int
		idx = -1
		for i, s := range dataSampah {
			if s.ID == ID {
				idx = i
				break
			}
		}
		if idx == -1 {
			fmt.Println("Nomor ID tidak ditemukan")
		}
		dataSampah = append(dataSampah[:idx], dataSampah[idx+1:]...)
		fmt.Println("Data berhasil dihapus")

		if konfirmasi() {
			break
		}
	}
}

func cariData() {

	for {
		if len(dataSampah) == 0 {
			fmt.Println("Data masih kosong.")
			return
		}

		var pilihan int
		var input string
		var inputjumlah int

		fmt.Println("Pilih metode pencarian:")
		fmt.Println("1. Berdasarkan jenis (Sequential search)")
		fmt.Println("2. Berdasarkan jumlah (Binary search)")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Print("Masukkan jenis sampah yang dicari: ")
			fmt.Scan(&input)
			index := sequentialSearch(input)
			if index != -1 {
				fmt.Printf("Data ditemukan di ID %d: %s\n", dataSampah[index].ID, dataSampah[index].Jenis)
				fmt.Printf("Jumlah: %d\n", dataSampah[index].Jumlah)
				fmt.Printf("Metode: %s\n", dataSampah[index].MetodeDaurUlang)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		} else if pilihan == 2 {
			fmt.Print("Masukkan jumlah sampah yang dicari: ")
			fmt.Scan(&inputjumlah)

			selectionSortByJumlah()
			index := binarySearch(inputjumlah)
			if index != -1 {
				fmt.Printf("Data ditemukan di ID %d: %s\n", dataSampah[index].ID, dataSampah[index].Jenis)
				fmt.Printf("Jumlah: %d\n", dataSampah[index].Jumlah)
				fmt.Printf("Metode: %s\n", dataSampah[index].MetodeDaurUlang)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		}

		if konfirmasi() {
			break
		}
	}
}

func sequentialSearch(jenis string) int {
	for i, s := range dataSampah {
		if strings.EqualFold(s.Jenis, jenis) {
			return i
		}
	}
	return -1
}

func binarySearch(jumlah int) int {
	insertionSortByJenis()

	kiri := 0
	kanan := len(dataSampah) - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if dataSampah[tengah].Jumlah == jumlah {
			return tengah
		} else if dataSampah[tengah].Jumlah < jumlah {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func urutkanData() {
	for {
		var pilihan int
		fmt.Println("Pilih metode pengurutan:")
		fmt.Println("1. Berdasarkan jumlah (Selection Sort)")
		fmt.Println("2. Berdasarkan jenis (Insertion Sort)")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			selectionSortByJumlah()
			fmt.Println("Data berhasil diurutkan berdasarkan jumlah.")
		case 2:
			insertionSortByJenis()
			fmt.Println("Data berhasil diurutkan berdasarkan jenis.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		tampilkanData()
		if konfirmasi() {
			break
		}
	}
}

func selectionSortByJumlah() {
	n := len(dataSampah)
	for i := 0; i < n-1; i++ {
		minIndeks := i
		for j := i + 1; j < n; j++ {
			if dataSampah[j].Jumlah < dataSampah[minIndeks].Jumlah {
				minIndeks = j
			}
		}
		dataSampah[i], dataSampah[minIndeks] = dataSampah[minIndeks], dataSampah[i]
	}
}

func insertionSortByJenis() {
	n := len(dataSampah)
	for i := 1; i < n; i++ {
		key := dataSampah[i]
		j := i - 1
		for j >= 0 && dataSampah[j].Jenis > key.Jenis {
			dataSampah[j+1] = dataSampah[j]
			j--
		}
		dataSampah[j+1] = key
	}
}

func tampilkanStatistik() {
	for {
		total := 0
		totalDaurUlang := 0

		for _, s := range dataSampah {
			total += s.Jumlah
			if s.MetodeDaurUlang != "" {
				totalDaurUlang += s.Jumlah
			}
		}

		fmt.Println("Total Sampah Terkumpul:", total)
		fmt.Println("Total Sampah Daur Ulang:", totalDaurUlang)

		if konfirmasi() {
			break
		}
	}
}

func tampilkan() {
	for {
		if len(dataSampah) == 0 {
			fmt.Println("Belum ada data sampah.")
			return
		}
		for _, s := range dataSampah {
			fmt.Printf("ID: %d. Jenis: %s | Jumlah: %d | Metode Daur Ulang: %s\n", s.ID, s.Jenis, s.Jumlah, s.MetodeDaurUlang)
		}
		if konfirmasi() {
			break
		}
	}
}

func tampilkanData() {

	if len(dataSampah) == 0 {
		fmt.Println("Belum ada data sampah.")
		return
	}
	for _, s := range dataSampah {
		fmt.Printf("ID: %d. Jenis: %s | Jumlah: %d | Metode Daur Ulang: %s\n", s.ID, s.Jenis, s.Jumlah, s.MetodeDaurUlang)
	}

}

func konfirmasi() bool {
	var konfirmasi string

	fmt.Print("Apakah ingin kembali ke menu utama? (ya/tidak) :")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "ya" {
		return true
	}
	return false
}
