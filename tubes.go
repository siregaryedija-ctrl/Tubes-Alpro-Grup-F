package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ============================================================
// TIPE DATA BENTUKAN
// ============================================================

// Kandidat menyimpan informasi lengkap satu kandidat dalam pemilihan
type Kandidat struct {
	NomorUrut int
	Nama      string
	Visi      string
	Misi      string
	PoinSuara int
}

// ============================================================
// VARIABEL GLOBAL
// Hanya array utama yang boleh menjadi variabel global.
// Jumlah data aktif (n) dan variabel lain dikelola sebagai
// parameter subprogram atau variabel lokal di main.
// ============================================================
var dataKandidat [100]Kandidat

// scanner adalah fasilitas baca stdin (infrastruktur I/O, bukan data bisnis)
var scanner = bufio.NewScanner(os.Stdin)

// ============================================================
// SUBPROGRAM UTILITAS I/O
// ============================================================

// readInput membaca satu baris teks dari stdin
// Output : string – teks yang dimasukkan pengguna
func readInput() string {
	scanner.Scan()
	return scanner.Text()
}

// readInt membaca bilangan bulat dari stdin.
// Jika input bukan angka, pengguna diminta mengulang (pola repeat-until).
// Output : int – nilai integer yang valid dari pengguna
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

// ============================================================
// SUBPROGRAM PENCARIAN
// ============================================================

// sequentialSearch mencari kandidat berdasarkan nomor urut secara berurutan
// dari elemen pertama hingga ditemukan atau habis.
// Input  : data [100]Kandidat – salinan array data kandidat
//
//	n    int           – jumlah data aktif
//	no   int           – nomor urut yang dicari
//
// Output : int  – indeks kandidat jika ditemukan, -1 jika tidak ditemukan
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

// binarySearch mencari kandidat berdasarkan nomor urut dengan metode biner.
// Prasyarat : data HARUS sudah terurut ascending berdasarkan NomorUrut.
// Input  : data [100]Kandidat – salinan array yang sudah terurut
//
//	n    int           – jumlah data aktif
//	no   int           – nomor urut yang dicari
//
// Output : int  – indeks kandidat jika ditemukan, -1 jika tidak ditemukan
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

// ============================================================
// SUBPROGRAM PENGURUTAN
// ============================================================

// selectionSort mengurutkan kandidat berdasarkan PoinSuara terbanyak (descending).
// Algoritma: setiap iterasi pilih elemen terbesar dari sisa, tukar ke posisi i.
// Input/Output : data *[100]Kandidat – array dimodifikasi langsung via pointer
//
//	n    int            – jumlah data aktif
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

// insertionSort mengurutkan kandidat berdasarkan NomorUrut secara ascending.
// Algoritma: setiap elemen disisipkan ke posisi yang tepat pada bagian terurut.
// Input/Output : data *[100]Kandidat – array dimodifikasi langsung via pointer
//
//	n    int            – jumlah data aktif
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
				// Jika sudah tidak perlu digeser, hentikan perulangan internal
				break
			}
		}

		data[j+1] = key
	}
	fmt.Printf("Data berhasil diurutkan")
}

// ============================================================
// SUBPROGRAM TAMPILAN
// ============================================================

// tampilProfil menampilkan informasi lengkap satu kandidat ke layar.
// Input : k Kandidat – data kandidat yang akan ditampilkan
func tampilProfil(k Kandidat) {
	fmt.Println("  --------------------------------")
	fmt.Println("  Nomor Urut :", k.NomorUrut)
	fmt.Println("  Nama       :", k.Nama)
	fmt.Println("  Visi 	  :", k.Visi)
	fmt.Println("  Misi		  :", k.Misi)
	fmt.Println("  Poin Suara :", k.PoinSuara)
	fmt.Println("  --------------------------------")
}

// hitungTotalSuara menjumlahkan seluruh poin suara dari semua kandidat aktif.
// Input  : data [100]Kandidat – salinan array data kandidat
//
//	n    int           – jumlah data aktif
//
// Output : int  – total suara seluruh kandidat
func hitungTotalSuara(data [100]Kandidat, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += data[i].PoinSuara
	}
	return total
}

// tampilkanSemuaKandidat mencetak profil semua kandidat yang terdaftar.
// Input : data [100]Kandidat – salinan array data kandidat
//
//	n    int           – jumlah data aktif
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

// tampilkanStatistik menampilkan statistik persentase suara tiap kandidat
// beserta total suara yang telah masuk.
// Input : data [100]Kandidat – salinan array data kandidat
//
//	n    int           – jumlah data aktif
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

// ============================================================
// SUBPROGRAM PEMBANTU PEMILIHAN METODE
// ============================================================

// pilihMetodeCari menampilkan sub-menu metode pencarian dan membaca pilihan pengguna.
// Pengguna diminta mengulang jika pilihan di luar 1-2 (pola repeat-until).
// Output : int – 1 untuk Sequential Search, 2 untuk Binary Search
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

// cariDenganMetode mencari kandidat menggunakan metode yang dipilih pengguna.
// Jika Binary Search dipilih, data diurutkan dahulu dengan Insertion Sort
// agar prasyarat binary search terpenuhi.
// Input  : data   *[100]Kandidat – array via pointer (bisa diurutkan di tempat)
//
//	n      int            – jumlah data aktif
//	no     int            – nomor urut yang dicari
//	metode int            – 1=Sequential Search, 2=Binary Search
//
// Output : int – indeks kandidat jika ditemukan, -1 jika tidak ditemukan
func cariDenganMetode(data *[100]Kandidat, n int, no int, metode int) int {
	if metode == 2 {
		fmt.Println("  [Binary Search] Mengurutkan data terlebih dahulu...")
		insertionSort(data, n, true)
		return binarySearch(*data, n, no)
	}
	return sequentialSearch(*data, n, no)
}

// ============================================================
// SUBPROGRAM OPERASI CRUD
// ============================================================

// tambahKandidat menambahkan satu kandidat baru ke akhir array.
// Menolak jika array sudah penuh (n >= 100).
// Input/Output : data *[100]Kandidat – array dimodifikasi via pointer
//
//	n    *int           – jumlah data aktif, bertambah 1 jika berhasil
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

// ubahKandidat mengubah nama dan visi-misi kandidat yang ditemukan.
// Pengguna memilih metode pencarian (Sequential atau Binary Search).
// Jika Binary Search dipilih, urutan array dapat berubah (data diurutkan dahulu).
// Input/Output : data *[100]Kandidat – array dimodifikasi via pointer
//
//	n    int            – jumlah data aktif (tidak berubah)
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

// hapusKandidat menghapus satu kandidat dari array berdasarkan nomor urut.
// Elemen dihapus dengan menggeser semua elemen sesudahnya satu posisi ke kiri,
// kemudian slot terakhir dikosongkan dan n dikurangi 1.
// Pengguna memilih metode pencarian (Sequential atau Binary Search).
// Input/Output : data *[100]Kandidat – array dimodifikasi via pointer
//
//	n    *int           – jumlah data aktif, berkurang 1 jika berhasil
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
		// Geser semua elemen di kanan idx satu posisi ke kiri
		for i := idx; i < *n-1; i++ {
			data[i] = data[i+1]
		}
		data[*n-1] = Kandidat{} // kosongkan slot terakhir
		*n--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Kandidat tidak ditemukan.")
	}
}

// tambahSuara menambahkan satu poin suara ke kandidat yang dipilih pemilih.
// Pencarian menggunakan Sequential Search (tidak mengubah urutan array).
// Input/Output : data *[100]Kandidat – array dimodifikasi via pointer
//
//	n    int            – jumlah data aktif
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

// ============================================================
// SUBPROGRAM MENU PENCARIAN & PENGURUTAN
// ============================================================

// menuSearch menjalankan pencarian kandidat
// dan menampilkan hasilnya.
// Input : data [100]Kandidat – salinan array (tidak mengubah urutan asli)
//
//	n    int           – jumlah data aktif
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

func menuSort(data *[100]Kandidat, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kandidat untuk diurutkan.")
		return
	}

	// 1. Dapatkan pilihan berdasarkan Suara atau Nomor
	berdasarkan := pilihSuaraNomor()

	// 2. Dapatkan pilihan arah urutan (Ascend/Descend)
	pilihanArah := pilihAscending()
	isAscending := true
	if pilihanArah == 2 {
		isAscending = false
	}

	// 3. Eksekusi algoritma yang tepat
	if berdasarkan == 1 {
		// Berdasarkan Suara -> Menggunakan Selection Sort milik Anda
		selectionSort(data, n, isAscending)
	} else if berdasarkan == 2 {
		// Berdasarkan Nomor -> Menggunakan Insertion Sort milik Anda
		insertionSort(data, n, isAscending)
	}

	fmt.Println("!") // Melengkapi teks "Data berhasil diurutkan" dari dalam fungsi sort
}

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

// ============================================================
// SUBPROGRAM MENU UTAMA
// ============================================================

// tampilMenu menampilkan pilihan menu utama dan membaca pilihan pengguna.
// Output : int – nomor menu yang dipilih pengguna
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

// ============================================================
// PROGRAM UTAMA
// ============================================================

func main() {
	// Perbesar buffer scanner untuk menampung input panjang (visi-misi)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	n := 0 // jumlah kandidat aktif saat ini (lokal di main)
	selesai := false

	// Pola repeat-until: ulangi tampil menu hingga pengguna memilih keluar
	for !selesai {
		pilihan := tampilMenu()
		switch pilihan {
		case 1:
			tambahKandidat(&dataKandidat, &n)
		case 2:
			ubahKandidat(&dataKandidat, n)
		case 3:
			hapusKandidat(&dataKandidat, &n)
		case 4: //searching
			menuSearch(dataKandidat, n)
		case 5: //sorting
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
