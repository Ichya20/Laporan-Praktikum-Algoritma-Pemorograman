package main

import "fmt"

func main() {
    var jumlahOrang int
    fmt.Scan(&jumlahOrang)
    var jumlahMotor int
    if jumlahOrang%2 == 0 {
        jumlahMotor = jumlahOrang / 2
    } else {
        jumlahMotor = (jumlahOrang / 2) + 1
    }
    fmt.Print(jumlahMotor)
}
