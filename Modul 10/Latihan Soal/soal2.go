package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Scan(&nam)
	if nam < 0 || nam > 100 {
		fmt.Println("Nilai tidak valid! Harus dalam rentang 0-100.")
		return
	}
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}
	fmt.Println(nmk)
}
