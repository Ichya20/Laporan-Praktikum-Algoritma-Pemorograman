package main

import (
	"fmt"
)

func main() {
	var totalBelanja, diskon int
	fmt.Scan(&totalBelanja)
	fmt.Scan(&diskon)
	totalSetelahDiskon := float64(totalBelanja) * (1 - float64(diskon)/100)
	fmt.Printf("%.2f\n", totalSetelahDiskon)
}
