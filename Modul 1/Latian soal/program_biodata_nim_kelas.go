package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membuat reader untuk membaca input pengguna
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Anda: ")
	nama, _ := reader.ReadString('\n')

	fmt.Print("Masukkan NIM Anda: ")
	nim, _ := reader.ReadString('\n')

	fmt.Print("Masukkan Kelas Anda: ")
	kelas, _ := reader.ReadString('\n')

	// Menampilkan resume singkat
	fmt.Println("\n--- Resume Singkat ---")
	fmt.Printf("Halo, nama saya %s", nama)
	fmt.Printf("Saya adalah mahasiswa dengan NIM %s", nim)
	fmt.Printf("Saya berasal dari kelas %s", kelas)
	fmt.Println("Saya sedang bersemangat untuk menyelesaikan studi saya dengan sukses dan mencapai tujuan akademis saya!")
}
