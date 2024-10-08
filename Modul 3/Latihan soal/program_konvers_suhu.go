package main

import (
	"fmt"
)

func main() {
	var celsius float64
	fmt.Print("Masukkan nilai temperatur dalam Celsius: ")
	fmt.Scan(&celsius)
	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	kelvin := (fahrenheit + 459.67) * 5 / 9
	fmt.Printf("Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Reamur: %.2f\n", reamur)
	fmt.Printf("Kelvin: %.2f\n", kelvin)
}
