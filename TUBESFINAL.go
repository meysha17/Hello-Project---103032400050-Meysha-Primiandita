package main

import "fmt"

// note oleh Meysha Primiandita - 103032400050

type Klub struct {
	Nama         string
	Main         int
	Menang       int
	Seri         int
	Kalah        int
	GolMasuk     int
	GolKemasukan int
	SelisihGol   int
	Poin         int
}

type Jadwal struct {
	Klub1 string
	Klub2 string
	Pekan int
}

const NMAX = 20

type Liga [NMAX]Klub
type DaftarJadwal [1000]Jadwal

func cariKlub(liga Liga, nama string) int { //SUDAH
	var i int

	for i = 0; i < NMAX; i++ {
		if liga[i].Nama == nama {
			return i
		}
	}
	return -1
}

func tambahKlub(liga *Liga) { //SUDAH
	var nama string
	var i, n, j, slotKosong int

	slotKosong = 0
	for i = 0; i < NMAX; i++ {
		if liga[i].Nama == "" {
			slotKosong++
		}
	}

	fmt.Print("Masukkan jumlah klub yang ingin diinput: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah klub harus lebih dari 0.")
	} else if n > slotKosong {
		fmt.Println("Input melebihi kuota klub yang tersedia. Silahkan kembali dan input ulang.")
	} else {
		i = 0
		for i < n {
			fmt.Print("Masukkan nama klub (3 huruf): ")
			fmt.Scan(&nama)
			if len(nama) != 3 {
				fmt.Println("Nama klub harus terdiri dari 3 huruf.")
			} else if cariKlub(*liga, nama) != -1 {
				fmt.Println("Klub sudah terdaftar.")
			} else {
				j = 0
				for j < NMAX && liga[j].Nama != "" {
					j++
				}
				if j < NMAX {
					liga[j].Nama = nama
					fmt.Println("Klub berhasil ditambahkan.")
					i++
				}
			}
		}
	}
}

func ubahKlub(liga *Liga) { //SUDAH
	var namaLama, namaBaru string
	var idx int

	fmt.Print("Masukkan nama klub yang ingin diubah: ")
	fmt.Scan(&namaLama)
	fmt.Print("Masukkan nama baru (3 huruf): ")
	fmt.Scan(&namaBaru)

	idx = cariKlub(*liga, namaLama)
	if idx != -1 {
		liga[idx].Nama = namaBaru
		fmt.Println("Nama klub berhasil diubah.")
	} else {
		fmt.Println("Klub tidak ditemukan.")
	}
}

func hapusKlub(liga *Liga) { //SUDAH
	var nama string
	var idx int

	fmt.Print("Masukkan nama klub yang ingin dihapus: ")
	fmt.Scan(&nama)

	idx = cariKlub(*liga, nama)
	if idx != -1 {
		liga[idx].Nama = ""
		fmt.Println("Klub berhasil dihapus.")
	} else {
		fmt.Println("Klub tidak ditemukan.")
	}
}

func buatJadwal(liga Liga, jadwal *DaftarJadwal, jumlahKlub int) int { //SUDAH
	var klubAktif [NMAX]string
	var k, i, idx, totalPekan, pekan int
	var last string

	k = 0
	for i = 0; i < NMAX; i++ {
		if liga[i].Nama != "" {
			klubAktif[k] = liga[i].Nama
			k++
		}
	}
	if jumlahKlub%2 != 0 {
		fmt.Println("Jumlah klub harus genap untuk membuat jadwal.")
		return 0
	}
	totalPekan = (jumlahKlub - 1) * 2
	idx = 0
	for pekan = 0; pekan < totalPekan/2; pekan++ {
		for i = 0; i < jumlahKlub/2; i++ {
			jadwal[idx] = Jadwal{Klub1: klubAktif[i], Klub2: klubAktif[jumlahKlub-1-i], Pekan: pekan + 1}
			idx++
		}
		last = klubAktif[jumlahKlub-1]
		for i = jumlahKlub - 1; i > 1; i-- {
			klubAktif[i] = klubAktif[i-1]
		}
		klubAktif[1] = last
	}
	for pekan = 0; pekan < totalPekan/2; pekan++ {
		for i = 0; i < jumlahKlub/2; i++ {
			jadwal[idx] = Jadwal{Klub1: klubAktif[jumlahKlub-1-i], Klub2: klubAktif[i], Pekan: pekan + 1 + totalPekan/2}
			idx++
		}
		last = klubAktif[jumlahKlub-1]
		for i = jumlahKlub - 1; i > 1; i-- {
			klubAktif[i] = klubAktif[i-1]
		}
		klubAktif[1] = last
	}
	return idx
}

func tampilkanJadwal(jadwal DaftarJadwal, jumlah int) { //SUDAH
	var i int

	for i = 0; i < jumlah; i++ {
		fmt.Printf("Pekan %d: %s vs %s\n", jadwal[i].Pekan, jadwal[i].Klub1, jadwal[i].Klub2)
	}
}

func inputHasil(liga *Liga) bool { //SUDAH
	var klub1, klub2 string
	var skor1, skor2 int
	var kandang, tandang int

	fmt.Print("Masukkan nama klub 1: ")
	fmt.Scan(&klub1)
	fmt.Print("Masukkan nama klub 2: ")
	fmt.Scan(&klub2)

	if klub1 == klub2 {
		fmt.Println("Tidak bisa input pertandingan untuk klub yang sama.")
		return false
	}

	fmt.Print("Masukkan skor klub 1: ")
	fmt.Scan(&skor1)
	fmt.Print("Masukkan skor klub 2: ")
	fmt.Scan(&skor2)

	if skor1 < 0 || skor2 < 0 {
		fmt.Println("Skor tidak boleh negatif.")
		return false
	}

	kandang = cariKlub(*liga, klub1)
	tandang = cariKlub(*liga, klub2)

	if kandang == -1 || tandang == -1 {
		fmt.Println("Salah satu atau kedua klub tidak ditemukan.")
		return false
	}

	liga[kandang].Main++
	liga[tandang].Main++
	liga[kandang].GolMasuk += skor1
	liga[kandang].GolKemasukan += skor2
	liga[tandang].GolMasuk += skor2
	liga[tandang].GolKemasukan += skor1

	if skor1 > skor2 {
		liga[kandang].Menang++
		liga[tandang].Kalah++
		liga[kandang].Poin += 3
	} else if skor2 > skor1 {
		liga[tandang].Menang++
		liga[kandang].Kalah++
		liga[tandang].Poin += 3
	} else {
		liga[kandang].Seri++
		liga[tandang].Seri++
		liga[kandang].Poin++
		liga[tandang].Poin++
	}

	liga[kandang].SelisihGol = liga[kandang].GolMasuk - liga[kandang].GolKemasukan
	liga[tandang].SelisihGol = liga[tandang].GolMasuk - liga[tandang].GolKemasukan
	return true
}

func hitungKlub(liga Liga) int { //SUDAH
	var sum int = 0
	var i int

	for i = 0; i < NMAX; i++ {
		if liga[i].Nama != "" {
			sum++
		}
	}
	return sum
}

func tampilkanPeringkat(liga Liga) { //SUDAH
	var dikunjungi [NMAX]bool
	var total int
	var urutan, i, idxmax int

	total = hitungKlub(liga)

	for urutan = 1; urutan <= total; urutan++ {
		idxmax = -1
		for i = 0; i < NMAX; i++ {
			if liga[i].Nama != "" && !dikunjungi[i] {
				if idxmax == -1 || liga[i].Poin > liga[idxmax].Poin ||
					(liga[i].Poin == liga[idxmax].Poin && liga[i].SelisihGol > liga[idxmax].SelisihGol) {
					idxmax = i
				}
			}
		}

		if idxmax != -1 {
			fmt.Printf("%2d. %-5s | Poin: %3d | SG: %3d\n", urutan, liga[idxmax].Nama, liga[idxmax].Poin, liga[idxmax].SelisihGol)
			dikunjungi[idxmax] = true
		}
	}
}

func sortLigaByPoin(liga *Liga, total int) { //SUDAH
	var i, j int
	var temp Klub

	for i = 1; i < total; i++ {
		temp = liga[i]
		j = i - 1
		for j >= 0 && (liga[j].Poin < temp.Poin ||
			(liga[j].Poin == temp.Poin && liga[j].SelisihGol < temp.SelisihGol)) {
			liga[j+1] = liga[j]
			j--
		}
		liga[j+1] = temp
	}
}

func binarySearchByPoin(liga Liga, total, carPoin int) int { //SUDAH
	var low, high, mid int
	low = 0
	high = total - 1

	for low <= high {
		mid = (low + high) / 2
		if liga[mid].Poin == carPoin {
			return mid
		} else if liga[mid].Poin < carPoin {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func cariByPoin(liga Liga) { //SUDAH
	var poinCari int
	var total int
	var i, idx int

	total = hitungKlub(liga)

	if total == 0 {
		fmt.Println("Belum ada klub yang terdaftar.")
	} else {
		fmt.Print("Masukkan jumlah poin yang ingin dicari: ")
		fmt.Scan(&poinCari)

		if poinCari < 0 {
			fmt.Println("Poin tidak boleh negatif.")
		} else {
			sortLigaByPoin(&liga, total)
			idx = binarySearchByPoin(liga, total, poinCari)

			if idx == -1 {
				fmt.Println("Tidak ada klub dengan poin tersebut.")
			} else {
				fmt.Printf("\nKlub dengan %d poin:\n", poinCari)

				i = idx
				for i >= 0 && liga[i].Poin == poinCari {
					i--
				}
				i++
				for i < total && liga[i].Poin == poinCari {
					fmt.Printf("- %s | Main: %d | Menang: %d | Seri: %d | Kalah: %d | GM: %d | GK: %d | SG: %d\n", liga[i].Nama, liga[i].Main, liga[i].Menang, liga[i].Seri, liga[i].Kalah, liga[i].GolMasuk, liga[i].GolKemasukan, liga[i].SelisihGol)
					i++
				}
			}
		}
	}
}

func main() {
	var liga Liga
	var jadwal DaftarJadwal
	var jumlahJadwal, jumlahKlub int
	var menu string
	var program bool = true

	for program {
		fmt.Println("\n--- MENU EPL MANAGER ---")
		fmt.Println("1. Tambah Klub")
		fmt.Println("2. Ubah Klub")
		fmt.Println("3. Hapus Klub")
		fmt.Println("4. Buat Jadwal Pertandingan")
		fmt.Println("5. Lihat Jadwal")
		fmt.Println("6. Input Hasil Pertandingan")
		fmt.Println("7. Lihat Peringkat")
		fmt.Println("8. Cari Berdasarkan Poin")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&menu)

		switch menu {
		case "1":
			tambahKlub(&liga)
		case "2":
			ubahKlub(&liga)
		case "3":
			hapusKlub(&liga)
		case "4":
			jumlahKlub = hitungKlub(liga)
			if jumlahKlub == 0 || jumlahKlub%2 != 0 {
				fmt.Println("Klub tidak cukup (harus genap & lebih dari 0)")
			} else {
				jumlahJadwal = buatJadwal(liga, &jadwal, jumlahKlub)
				fmt.Println("Jadwal berhasil dibuat.")
			}
		case "5":
			jumlahKlub = hitungKlub(liga)

			if jumlahKlub == 0 {
				fmt.Println("Belum ada klub yang terdaftar.")
			} else if jumlahKlub%2 != 0 {
				fmt.Println("Jumlah klub ganjil, jadwal tidak bisa dibuat. Harap tambah atau hapus klub agar genap.")
			} else if jumlahJadwal == 0 {
				fmt.Println("Jadwal belum dibuat.")
			} else {
				tampilkanJadwal(jadwal, jumlahJadwal)
			}
		case "6":
			inputHasil(&liga)
		case "7":
			tampilkanPeringkat(liga)
		case "8":
			cariByPoin(liga)
		case "0":
			fmt.Println("Terima kasih telah menggunakan EPL Manager")
			program = false
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}
