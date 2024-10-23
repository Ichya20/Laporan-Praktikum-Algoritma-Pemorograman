package main

import "fmt"

func main() {
    var bilangan, hasil int
    fmt.Scan(&bilangan)
    hasil = 1
    for i := 1; i <= bilangan; i++ {
        hasil *= i
    }
    fmt.Println(hasil)
}
