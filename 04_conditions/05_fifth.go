package main

import (
    "fmt"
)

func whichIsBigger(x, y int) int {
    if x == y {
        return -1
    } else if x > y {
        return x
    }
    return y
}

func main() {
    result := whichIsBigger(10, 10)
    if result != -1 {
        fmt.Println(result)
    } else {
        fmt.Println("Are equal")
    }
}