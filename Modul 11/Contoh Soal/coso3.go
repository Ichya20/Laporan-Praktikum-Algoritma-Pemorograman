package main

import "fmt"

func main() {
	var tipe_kendaraan string
	var durasi, tarif int

	fmt.Println("Masukkan jenis kendaraan (Motor/Mobil/Truk):")
	fmt.Scan(&tipe_kendaraan)
	fmt.Println("Masukkan durasi parkir (dalam jam) :")
	fmt.Scan(&durasi)

	switch {
	case tipe_kendaraan == "Motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case tipe_kendaraan == "Motor" && durasi >= 2:
		tarif = 9000
	case tipe_kendaraan == "Mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case tipe_kendaraan == "Mobil" && durasi >= 2:
		tarif = 20000
	case tipe_kendaraan == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case tipe_kendaraan == "Truk" && durasi >= 2:
		tarif = 35000
	default:
		fmt.Println("Tipe kendaraan tidak dikenal")
	}
	fmt.Printf("Tarif Parkir: Rp. %d", tarif)
}