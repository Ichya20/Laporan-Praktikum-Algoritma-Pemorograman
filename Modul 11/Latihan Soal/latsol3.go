package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	switch {
	case bilangan%2 != 0:
		fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, bilangan+(bilangan+1))
	case bilangan%2 == 0 && bilangan%10 != 0:
		fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, bilangan*(bilangan+1))
	case bilangan%5 == 0 && bilangan%10 != 0:
		fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d ^2 = %d\n", bilangan, bilangan*bilangan)
	case bilangan%10 == 0:
		fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", bilangan, bilangan/10)
	default:
		fmt.Println("Kategori tidak ditemukan.")
	}
}
