package main

import (
    "fmt"
)

var (
    memoryAddress = "01001100"
    c byte
)

func main() {
    c = 'A'
    fmt.Printf("%T %v", memoryAddress, c)
}
