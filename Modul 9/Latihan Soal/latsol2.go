package main

import "fmt"

func main() {
    var number int
    fmt.Scan(&number)
    if number < 0 && number%2 == 0 {
        fmt.Println("genap negatif")
    } else {
        fmt.Println("bukan")
    }
}
