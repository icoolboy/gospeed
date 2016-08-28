package main

import (
    "fmt"
)

func sum2number(x, y int) int {
    return x + y
}

func main() {
    total := sum2number(10, 2)
    fmt.Printf("Total of sum : %d", total)
}