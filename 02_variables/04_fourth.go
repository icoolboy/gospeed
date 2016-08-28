package main

import (
    "fmt"
)

var (
    memoryAddress = "01001100"
    c byte
)

func main() {
    year := 2016
    month := "July"
    fmt.Printf("%T %T", year, month)
    fmt.Println(year, month)
}
