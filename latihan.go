package main

import "fmt"

func fibonacciGanjil(n int) {
    a, b := 0, 1
    i := 0
    for i < n {
        a, b = b, a+b
        if a%2 != 0 {
            fmt.Print(a, " ")
            i++
        }
    }
    fmt.Println()
}
func main() {
    var n int
    fmt.Print("Masukkan jumlah bilangan Fibonacci: ")
    fmt.Scan(&n)
    fmt.Println("Deret Fibonacci (hanya ganjil):")
    fibonacciGanjil(n)
}
