package main

import "fmt"

func main() {
	var beratParsel int
	var totalKg, sisaGram int
	var biayaKg, biayaSisa, totalBiaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)
	totalKg = beratParsel / 1000
	sisaGram = beratParsel % 1000
	biayaKg = totalKg * 10000
	if totalKg > 10 {
		biayaSisa = 0
	} else if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else {
		biayaSisa = sisaGram * 15
	}
	totalBiaya = biayaKg + biayaSisa
	fmt.Printf("Detail berat: %d kg + %d gr\n", totalKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
