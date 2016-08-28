package main

import (
    "fmt"
)

var country = "IRAN"

func main() {
    var (
        number int8 = -10
        myFirstName string = "Erfan"
        check  bool
    )
    
    fmt.Printf("%d %s", number, myFirstName)
    fmt.Printf("%b", check)
    fmt.Println(country)
}