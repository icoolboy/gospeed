package main

import (
    "fmt"
)

//Closures function
func sum2number(i int) func() int {
    return func() int {
        return i
    }
}

func main() {
    total := sum2number(10)
    fmt.Printf("Total of sum : %d", total())
}