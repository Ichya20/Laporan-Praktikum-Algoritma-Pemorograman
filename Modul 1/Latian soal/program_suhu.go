package main

import "fmt"

func main() {
	// Deklarasi variabel untuk suhu Fahrenheit
	var fahrenheit float64

	// Meminta input suhu dalam Fahrenheit
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Menghitung konversi ke Celsius
	celsius := (5.0 / 9.0) * (fahrenheit - 32)

	// Menampilkan hasil suhu dalam Celsius
	fmt.Printf("Suhu dalam Celsius: %.2f\n", celsius)
}
