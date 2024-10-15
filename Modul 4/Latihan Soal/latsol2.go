package main

import (
	"fmt"
)

func main() {
	var bmi, tinggi float64
	fmt.Scan(&bmi)
	fmt.Scan(&tinggi)
	berat := bmi * tinggi * tinggi
	fmt.Printf("%.2f\n", berat)
}
