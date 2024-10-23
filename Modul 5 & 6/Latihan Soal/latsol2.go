package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    const pi = 3.14159 
    for i := 1; i <= n; i++ {
        var r, t float64
        fmt.Scan(&r, &t)
        volume := (1.0 / 3.0) * pi * r * r * t 
        fmt.Println(volume)
    }
}
