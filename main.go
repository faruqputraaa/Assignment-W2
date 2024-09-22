package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Data pengguna disimpan dalam map
var dataPengguna = make(map[string]int)

// membersihkan terminal
const clear = "\033[2J\033[H"

func main() {
	for {
		// Menu utama
		fmt.Println("\n===== Menu Utama =====")
		fmt.Println("1. Tampilkan Hello World")
		fmt.Println("2. Operasi Matematika Sederhana")
		fmt.Println("3. Simpan dan Tampilkan Data Pengguna")
		fmt.Println("4. Hitung Faktorial (Rekursif)")
		fmt.Println("5. Hitung Total (Variadic)")
		fmt.Println("tekan 'exit' untuk keluar")

		// Input pilihan pengguna
		fmt.Print("Pilih menu: ")
		pilihan := bacaInput()

		// Switch untuk menangani pilihan menu
		switch pilihan {
		case "1":
			fmt.Println(clear)
			tampilkanHelloWorld()
		case "2":
			fmt.Println(clear)
			lakukanOperasiMatematika()
		case "3":
			fmt.Println(clear)
			kelolaDataPengguna()
		case "4":
			fmt.Println(clear)
			hitungFaktorial()
		case "5":
			fmt.Println(clear)
			hitungTotal()
		case "exit":
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			return
		}
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Tekan enter untuk kembali ke Menu...")
		fmt.Scanln()
		fmt.Println("\033[2J\033[H")
	}
}

// Fungsi untuk menampilkan "Hello, Dunia!"
func tampilkanHelloWorld() {
	fmt.Println("Hello World!")
}

// Fungsi untuk melakukan operasi matematika dasar
func lakukanOperasiMatematika() {
	fmt.Println("\n-- Operasi Matematika --")
	fmt.Print("Masukkan angka pertama: ")
	angka1 := bacaInputFloat()
	fmt.Print("Masukkan angka kedua: ")
	angka2 := bacaInputFloat()

	// Melakukan operasi penjumlahan
	hasilPenjumlahan := angka1 + angka2
	fmt.Printf("Hasil Penjumlahan: %.2f\n", hasilPenjumlahan)

	// Melakukan operasi pengurangan
	hasilPengurangan := angka1 - angka2
	fmt.Printf("Hasil Pengurangan: %.2f\n", hasilPengurangan)

	// Melakukan operasi perkalian
	hasilPerkalian := angka1 * angka2
	fmt.Printf("Hasil Perkalian: %.2f\n", hasilPerkalian)

	// Melakukan operasi pembagian
	if angka2 != 0 {
		hasilPembagian := angka1 / angka2
		fmt.Printf("Hasil Pembagian: %.2f\n", hasilPembagian)
	} else {
		fmt.Println("Tidak dapat melakukan pembagian dengan nol.")
	}

	// Melakukan operasi sisa bagi
	if angka2 != 0 {
		hasilSisaBagi := int(angka1) % int(angka2)
		fmt.Printf("Hasil Sisa Bagi: %d\n", hasilSisaBagi)
	} else {
		fmt.Println("Tidak dapat melakukan sisa bagi dengan nol.")
	}
}

// Fungsi untuk mengelola data pengguna dengan tambahan filter umur
func kelolaDataPengguna() {
	fmt.Println("\n-- Simpan dan Tampilkan Data Pengguna --")
	for {
		fmt.Println("1. Tambah Data Pengguna")
		fmt.Println("2. Tampilkan Semua Data")
		fmt.Println("3. Filter Data Pengguna")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Pilih opsi: ")
		pilihan := bacaInput()

		switch pilihan {
		case "1":
			fmt.Println(clear)
			tambahDataPengguna()
		case "2":
			fmt.Println(clear)
			tampilkanDataPengguna()
		case "3":
			fmt.Println(clear)
			filterDataPengguna()
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// Fungsi untuk memfilter data pengguna menggunakan closure
func filterDataPengguna() {
	fmt.Print("Masukkan umur minimum untuk filter: ")
	umurMinimum := bacaInputInt()

	// Memfilter data berdasarkan umur
	filter := func(data map[string]int, kondisi func(int) bool) map[string]int {
		hasil := make(map[string]int)
		for nama, umur := range data {
			if kondisi(umur) {
				hasil[nama] = umur
			}
		}
		return hasil
	}

	// Closure sebagai kondisi untuk filter umur
	kondisiUmur := func(umur int) bool {
		return umur > umurMinimum
	}

	// Memanggil fungsi filter dengan closure kondisiUmur
	hasilFilter := filter(dataPengguna, kondisiUmur)

	// Menampilkan hasil filter
	if len(hasilFilter) == 0 {
		fmt.Println("Tidak ada pengguna yang memenuhi kriteria umur.")
	} else {
		fmt.Println("\n-- Data Pengguna yang Difilter --")
		for nama, umur := range hasilFilter {
			fmt.Printf("Nama: %s, Umur: %d\n", nama, umur)
		}
	}
}

// Fungsi untuk menambah data pengguna ke map
func tambahDataPengguna() {
	fmt.Print("Masukkan nama pengguna: ")
	nama := bacaInput()
	fmt.Print("Masukkan umur pengguna: ")
	umur := bacaInputInt()

	dataPengguna[nama] = umur
	fmt.Println("Data berhasil ditambahkan.")
}

// Fungsi untuk menampilkan data pengguna
func tampilkanDataPengguna() {
	if len(dataPengguna) == 0 {
		fmt.Println("Tidak ada data pengguna yang tersimpan.")
		return
	}
	fmt.Println("\n-- Data Pengguna --")
	for nama, umur := range dataPengguna {
		fmt.Printf("Nama: %s, Umur: %d\n", nama, umur)
	}
}

// Fungsi rekursif untuk menghitung faktorial
func hitungFaktorial() {
	fmt.Print("Masukkan angka untuk menghitung faktorial: ")
	angka := bacaInputInt()
	fmt.Printf("Faktorial dari %d adalah %d\n", angka, faktorial(angka))
}

// Rekursif function faktorial
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

// Fungsi variadic untuk menghitung total
func hitungTotal() {
	fmt.Println("\n-- Hitung Total (Variadic) --")
	fmt.Print("Masukkan angka-angka (pisahkan dengan spasi): ")
	input := bacaInput()
	angkaStr := strings.Split(input, " ")
	angka := []int{}
	for _, s := range angkaStr {
		num, err := strconv.Atoi(s)
		if err == nil {
			angka = append(angka, num)
		}
	}
	fmt.Printf("Total: %d\n", total(angka...))
}

// Variadic function untuk menghitung total angka
func total(angka ...int) int {
	sum := 0
	for _, num := range angka {
		sum += num
	}
	return sum
}

// Fungsi untuk membaca input string dari pengguna
func bacaInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi untuk membaca input integer
func bacaInputInt() int {
	input := bacaInput()
	angka, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input tidak valid, silakan masukkan angka.")
		return bacaInputInt()
	}
	return angka
}

// Fungsi untuk membaca input float
func bacaInputFloat() float64 {
	input := bacaInput()
	angka, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input tidak valid, silakan masukkan angka.")
		return bacaInputFloat()
	}
	return angka
}
