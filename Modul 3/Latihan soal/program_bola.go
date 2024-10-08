package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&r)
	const pi = 3.1415926535
	volume := (4.0 / 3.0) * pi * math.Pow(r, 3)
	luas := 4 * pi * math.Pow(r, 2)
	fmt.Printf("Volume bola: %.2f\n", volume)
	fmt.Printf("Luas kulit bola: %.2f\n", luas)
}
