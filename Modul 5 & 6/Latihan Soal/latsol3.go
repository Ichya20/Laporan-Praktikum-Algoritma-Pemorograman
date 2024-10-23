package main

import "fmt"

func main() {
    var bilangan, pangkat, hasil int
    fmt.Scan(&bilangan)
    fmt.Scan(&pangkat)
    hasil = 1
    for i := 1; i <= pangkat; i++ {
        hasil *= bilangan
    }
    fmt.Println(hasil)
}
