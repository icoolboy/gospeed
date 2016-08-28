package main

import (
    "fmt"
)

func main() {
    x, y := 10, 1
    //Count of `if` in this sample : 1
    if x == y {
        fmt.Println("`x` is equal to `y` !")
    } else if x > y {
        fmt.Println("`x` is bigger than `y` !")
    } else {
        fmt.Printf("%s", "`x` is smaller than `y` :D")
    }
}