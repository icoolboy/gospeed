package main

import (
    "fmt"
)

func main() {
    s := make([]int, 0, 5)
    n := []string{"Hi", "Hello", "Welcome", "GoodBye"}
    fmt.Println(s)
    fmt.Println(len(s))
    fmt.Println(cap(s))
    
    fmt.Println("-------------------")
    
    fmt.Println(n)
    fmt.Println(len(n))
    fmt.Println(cap(n))
}