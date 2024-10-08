package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&year)
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		fmt.Println("true")  // Tahun kabisat
	} else {
		fmt.Println("false") // Bukan tahun kabisat
	}
}
