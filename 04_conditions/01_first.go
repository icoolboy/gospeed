package main

import (
    "fmt"
)

func main() {
    x, y := 10, 2
    if x > y {
        fmt.Println("`x` bigger than `y` !")
    } else {
        fmt.Printf("%s", "`y` bigger than `y` or equal to `y`")
    }
}