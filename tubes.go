package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Tipe Bentukan Kandidat, menyimpan informasi lengkap satu kandidat dalam pemilihan
type Kandidat struct {
	NomorUrut int
	Nama      string
	Visi      string
	Misi      string
	PoinSuara int
}

// Array Global yang digunakan untuk menyimpan array tiap tiap kandidat
var dataKandidat [100]Kandidat

// Variabel scanner digunakan untuk input 
// memanfaatkan package bufio (buffered I/O) yang bisa digunakan untuk membaca input satu satu baris di CLI
// os.stdin digunakan untuk mengarahkan alur data untuk disimpan di memori
var scanner = bufio.NewScanner(os.Stdin)

// fungsi readInput membaca satu baris teks (string)
// P.S. string – teks yang dimasukkan pengguna
func readInput() string {
	scanner.Scan()
	return scanner.Text()
}

// readInt membaca bilangan bulat (integer)
// P.S. integer angka yang diinputkan
// Jika input bukan angka, pengguna diminta mengulang (pola repeat-until).
func readInt() int {
	for {
		str := readInput()
		val, err := strconv.Atoi(str)
		if err == nil {
			return val
		}
		fmt.Print("Input tidak valid, masukkan angka: ")
	}
}

// sequentialSearch mencari kandidat berdasarkan nomor urut secara berurutan dari elemen pertama hingga ditemukan atau elemen terakhir array.
// I.S. data sebgai menampung array global semua kandidat, n sebagai jumlah kandidat, no sebagai nomor kandidat yang dicari
// P.S. -1 jika nomor kandidat tidak ditemukan, atau idx berisi nomor kandiat yang dicari
func sequentialSearch(data [100]Kandidat, n int, no int) int {
	idx := -1
	i := 0
	for i < n && idx == -1 {
		if data[i].NomorUrut == no {
			idx = i
		}
		i++
	}
	return idx
}

// binarySearch mencari kandidat berdasarkan nomor urut yang sudah diurutkan dengan metode biner (Devide and Conquer) 
// I.S. data sebgai menampung array global semua kandidat, n sebagai jumlah kandidat, no sebagai nomor kandidat yang dicari
// P.S. -1 jika nomor kandidat tidak ditemukan, atau idx berisi nomor kandiat yang dicari
func binarySearch(data [100]Kandidat, n int, no int) int {
	low := 0
	high := n - 1
	idx := -1
	for low <= high && idx == -1 {
		mid := low + (high-low)/2
		if data[mid].NomorUrut == no {
			idx = mid
		} else if data[mid].NomorUrut < no {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

// selectionSort mengurutkan kandidat berdasarkan PoinSuara (Pengurutan berdasarkan poin suara menggunakan selection sort)
// Algoritma: setiap iterasi pilih elemen terbesar dari sisa, tukar ke posisi i.
func selectionSort(data *[100]Kandidat, n int, isAscending bool) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending == true {
				if data[j].PoinSuara < data[maxIdx].PoinSuara {
					maxIdx = j
				}
			} else if isAscending == false {
				if data[j].PoinSuara > data[maxIdx].PoinSuara {
					maxIdx = j
				}
			}

		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
	fmt.Printf("Data berhasil diurutkan")
}

// insertionSort mengurutkan kandidat berdasarkan NomorUrut (Pengurutan berdasarkan Nomor urut kandiat menggunakan insrtion sort)
// Algoritma: setiap elemen disisipkan ke posisi yang tepat pada bagian terurut.
func insertionSort(data *[100]Kandidat, n int, isAscending bool) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 {
			harusDigeser := false

			if isAscending == true {
				if data[j].NomorUrut > key.NomorUrut {
					harusDigeser = true
				}
			} else if isAscending == false {
				if data[j].NomorUrut < key.NomorUrut {
					harusDigeser = true
				}
			}

			if harusDigeser {
				data[j+1] = data[j]
				j--
			} else {
				break
			}
		}

		data[j+1] = key
	}
	fmt.Printf("Data berhasil diurutkan")
}

// tampilProfil menampilkan informasi lengkap satu kandidat ke layar.
// I.S. k sebgai informasi lengkap satu kandidat
// P.S. menampilkan informasi satu kandidat
func tampilProfil(k Kandidat) {
	fmt.Println("  --------------------------------")
	fmt.Println("  Nomor Urut :", k.NomorUrut)
	fmt.Println("  Nama       :", k.Nama)
	fmt.Println("  Visi 	  :", k.Visi)
	fmt.Println("  Misi		  :", k.Misi)
	fmt.Println("  Poin Suara :", k.PoinSuara)
	fmt.Println("  --------------------------------")
}

//hintungTotalSuara menghitung total suara (vote) yang telah dilakukan
// I.S. data sebagai array global dataKandidat, n jumlah kandidat dalam pemilihan
// P.S. total berisi total suara dari elemen tipe bentukan PoinSuara semua kandidat 
func hitungTotalSuara(data [100]Kandidat, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += data[i].PoinSuara
	}
	return total
}

// tampilkanSemuaKandidat menampilkan semua profil kandidat
// I.S. data sebagai array global dataKandidat, n jumlah kandidat dalam pemilihan
// P.S. menampilkan semua profil kandidat lewat procedure tampilProfil
func tampilkanSemuaKandidat(data [100]Kandidat, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}
	fmt.Printf("\n=== Daftar %d Kandidat ===\n", n)
	for i := 0; i < n; i++ {
		tampilProfil(data[i])
	}
}

// Menampilkan Statistik perolehan suara dari semua kandidat
// P.S. Menampilkan perolehan suara (persentase dan jumlah suara) dari semmua kandidat
func tampilkanStatistik(data [100]Kandidat, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}
	total := hitungTotalSuara(data, n)
	fmt.Println("\n=== Statistik Pemilihan ===")
	fmt.Println("Total Suara Masuk:", total)
	fmt.Println()
	for i := 0; i < n; i++ {
		persentase := 0.0
		if total > 0 {
			persentase = float64(data[i].PoinSuara) / float64(total) * 100
		}
		fmt.Printf("  [%d] %-20s : %3d suara (%.2f%%)\n",
			data[i].NomorUrut, data[i].Nama, data[i].PoinSuara, persentase)
	}
}

// pilihMetodeCari menampilkan sub-menu metode pencarian dan membaca pilihan pengguna.
func pilihMetodeCari() int {
	fmt.Println("  Pilih metode pencarian:")
	fmt.Println("  1. Sequential Search")
	fmt.Println("  2. Binary Search")
	fmt.Print("  Pilihan: ")
	pilihan := 0
	for pilihan != 1 && pilihan != 2 {
		pilihan = readInt()
		if pilihan != 1 && pilihan != 2 {
			fmt.Print("  Pilihan tidak valid, masukkan 1 atau 2: ")
		}
	}
	return pilihan
}

// untuk memilih pengurutan secar ascending atau descending
func pilihAscending() int {
	fmt.Println("  Pilih metode pengurutan nomor:")
	fmt.Println("  1. Ascending")
	fmt.Println("  2. Descending")
	fmt.Print("  Pilihan: ")

	pilihan := 0

	for pilihan != 1 && pilihan != 2 {
		pilihan = readInt()

		if pilihan != 1 && pilihan != 2 {
			fmt.Print("  Pilihan tidak valid, masukkan 1 atau 2: ")
		}
	}

	return pilihan
}

// cariDenganMetode mencari kandidat menggunakan metode yang dipilih
// Jika Binary Search dipilih, data diurutkan dahulu dengan Insertion Sort agar prasyarat binary search terpenuhi.
// I.S. mendapatkan array global kandidiat (in/out)(untuk menjadlankan Binary Search), jumlah kandidat, nomor kandiat yang dipilih, dan metode pencarian
// P.S. mengembalikan  index array dari kandidat yang dipilih 
func cariDenganMetode(data *[100]Kandidat, n int, no int, metode int) int {
	if metode == 2 {
		fmt.Println("  [Binary Search] Mengurutkan data terlebih dahulu...")
		insertionSort(data, n, true)
		return binarySearch(*data, n, no)
	}
	return sequentialSearch(*data, n, no)
}

// tambahKandidat menambahkan satu kandidat baru ke akhir array global kandidat.
// Menolak jika array sudah penuh (n >= 100).
// I.S. array global kandidat untuk menampung kandidat baru (in/out) dan n jumlah kandidat yang bertambah (in/out)
// P.S. data kandidat yang diinputkan akan tersimpan dalam array bentukan Kandidat yang tersimpan dalam array global dataKandidat
func tambahKandidat(data *[100]Kandidat, n *int) {
	if *n >= 100 {
		fmt.Println("Data kandidat sudah penuh (maksimal 100 kandidat).")
		return
	}
	var k Kandidat
	fmt.Print("Masukkan Nomor Urut       : ")
	k.NomorUrut = readInt()

	idx := sequentialSearch(*data, *n, k.NomorUrut)
	if idx != -1 {
		fmt.Println("Nomor urut tersebut sudah terisi")
		return
	}
	if k.NomorUrut <= 0 {
		fmt.Println("Nomor urut tidak valid")
		return
	}

	fmt.Print("Masukkan Nama             : ")
	k.Nama = readInput()
	fmt.Print("Masukkan Visi	         : ")
	k.Visi = readInput()
	fmt.Print("Masukkan Misi             : ")
	k.Misi = readInput()
	k.PoinSuara = 0
	data[*n] = k
	*n++
	fmt.Println("Data berhasil ditambahkan.")
}

// Mengubah nama, visi, dan misi baru pada profil kandidat yang nomornya dipilih
// I.S. array global kandidat In/Out untuk menyimpan perubahann
// Profil data pada no kandidat yang dipilih berubah
func ubahKandidat(data *[100]Kandidat, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}
	fmt.Print("Masukkan Nomor Urut kandidat yang akan diubah: ")
	no := readInt()
	idx := sequentialSearch(*data, n, no)
	if idx != -1 {
		fmt.Println("Data saat ini:")
		tampilProfil(data[idx])
		fmt.Print("Masukkan Nama Baru        : ")
		data[idx].Nama = readInput()
		fmt.Print("Masukkan Visi 			 : ")
		data[idx].Visi = readInput()
		fmt.Print("Masukkan Misi Baru		 : ")
		data[idx].Misi = readInput()
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Kandidat tidak ditemukan.")
	}
}

// Menghapus kandidat yang dipilih
// I.S. array global semua kandidat In/Out dan n jumlah kandidat In/Out
// P.S. kandidat yang dipilih akan dihapus dari array global kandidat
func hapusKandidat(data *[100]Kandidat, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}
	fmt.Print("Masukkan Nomor Urut kandidat yang akan dihapus: ")
	no := readInt()
	idx := sequentialSearch(*data, *n, no)
	if idx != -1 {
		fmt.Println("Data yang akan dihapus:")
		tampilProfil(data[idx])
		// Menggeser semua elemen di kanan idx satu posisi ke kiri
		for i := idx; i < *n-1; i++ {
			data[i] = data[i+1]
		}
		data[*n-1] = Kandidat{}
		*n--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Kandidat tidak ditemukan.")
	}
}

// Untuk menambahkan suara atau voting pada kandidat yang dipilih
// I.S. data kandidat In/Out kareana PoinSuara akan menyimpan vote
// P.S. properti PoinSuara kandidat yang dipilih akan mengalami penambahan saru suara 
func tambahSuara(data *[100]Kandidat, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}
	fmt.Print("Masukkan Nomor Urut kandidat yang dipilih: ")
	no := readInt()
	idx := sequentialSearch(*data, n, no)
	if idx != -1 {
		data[idx].PoinSuara++
		fmt.Printf("Suara untuk [%d] %s berhasil ditambahkan.\n",
			data[idx].NomorUrut, data[idx].Nama)
	} else {
		fmt.Println("Kandidat tidak ditemukan.")
	}
}

// menuSearch menjalankan pencarian kandidat dan menampilkan hasilnya.
// I.S. mengambil array global kandidat dan jumlahkandidat
// P.S. menampilkan profil kandidat yang dicari
func menuSearch(data [100]Kandidat, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}
	
	fmt.Print("Masukkan Nomor Urut yang dicari: ")
	no := readInt()
	metode := pilihMetodeCari()
	idx := cariDenganMetode(&data, n, no, metode)

	if idx != -1 {
		fmt.Println("Kandidat ditemukan :")
		tampilProfil(data[idx])
	} else {
		fmt.Println("Kandidat tidak ditemukan.")
	}
}

// Menu pada fitur Urutkan (sortting data)
// I.S. mengambil array global kandidat dan jumlahkandidat
// P.S. data diurutkan berdasarkan pilihan pada menu (suara/nomo) dan (ascending/descending)
func menuSort(data *[100]Kandidat, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kandidat untuk diurutkan.")
		return
	}

	// menampilkan pilihan untuk pengurutan sesuaai suara atau nomor (kembalian int)
	berdasarkan := pilihSuaraNomor()

	// menampilkan pilihan untuk pengurutan dengan ascending atau descending (kembalian int)
	pilihanArah := pilihAscending()
	// asumsi awal ascending
	isAscending := true
	// jika pilihan menjadi decending (2) maka isAscending menjadi false
	if pilihanArah == 2 {
		isAscending = false
	}

	if berdasarkan == 1 {
		selectionSort(data, n, isAscending)
	} else if berdasarkan == 2 {
		insertionSort(data, n, isAscending)
	}

	fmt.Println("!")
}

// untuk menampilkan pilihan Suara atau Nomor pada menu Urutkan
// I.S. menmpilkan pilihan untuk memilih antara suara atau nommor
// P.S. mengembalikan pilihan untuk menu Urutkan
func pilihSuaraNomor() int {
	fmt.Println("  Pilih diurutkan berdasarkan :")
	fmt.Println("  1. Suara")
	fmt.Println("  2. Nomor")
	fmt.Print("  Pilihan: ")

	pilihan := 0

	for pilihan != 1 && pilihan != 2 {
		pilihan = readInt()

		if pilihan != 1 && pilihan != 2 {
			fmt.Print("  Pilihan tidak valid, masukkan 1 atau 2: ")
		}
	}
	return pilihan
}

// Procedure untuk menampilkan menu
func tampilMenu() int {
	fmt.Println("\n╔══════════════════════════════════════════════════╗")
	fmt.Println("║   Sistem Pemungutan Suara Digital (E-Voting)     ║")
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Println("║  1.  Tambah Kandidat                             ║")
	fmt.Println("║  2.  Ubah Kandidat                               ║")
	fmt.Println("║  3.  Hapus Kandidat                              ║")
	fmt.Println("║  4.  Cari Kandidat                               ║")
	fmt.Println("║  5.  Urutkan                                     ║")
	fmt.Println("║  6.  Tampilkan Semua Kandidat                    ║")
	fmt.Println("║  7.  Tampilkan Statistik                         ║")
	fmt.Println("║  8. Tambah Suara (Voting)                        ║")
	fmt.Println("║  0.  Keluar                                      ║")
	fmt.Println("╚══════════════════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
	return readInt()
}

func main() {
	// kondisi awal jumlah kandidat
	n := 0 
	// untuk mengatur perulangan menu
	selesai := false

	//perulangan untuk menu
	for !selesai {
		pilihan := tampilMenu()
		switch pilihan {
		case 1:
			tambahKandidat(&dataKandidat, &n)
		case 2:
			ubahKandidat(&dataKandidat, n)
		case 3:
			hapusKandidat(&dataKandidat, &n)
		case 4: 
			menuSearch(dataKandidat, n)
		case 5: 
			menuSort(&dataKandidat, n)
		case 6:
			tampilkanSemuaKandidat(dataKandidat, n)
		case 7:
			tampilkanStatistik(dataKandidat, n)
		case 8:
			tambahSuara(&dataKandidat, n)
		case 0:
			fmt.Println("Terima kasih. Program selesai.")
			selesai = true
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
