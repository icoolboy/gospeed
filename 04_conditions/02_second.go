package main

import (
    "fmt"
)

func main() {
    x, y := 10, 1
    //Count of `if` in this sample : 3
    if x == y {
        fmt.Println("`x` is equal to `y` !")
    }
    
    if x > y {
        fmt.Println("`x` is bigger than `y` !")
    }
    
    if x < y {
        fmt.Printf("%s", "`x` is smaller than `y` :D")
    }
}