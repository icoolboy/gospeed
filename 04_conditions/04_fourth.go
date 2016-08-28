package main

import (
    "fmt"
    "strings"
)

func main() {
    var (
        name = "Rob"
    )
    
    x, y := 10, 1
    
    //Usually, `switch` stmt used when we need high count of `if`
    
    switch {
        case x > y :
            fmt.Println("`x` is bigger than `y` !")
        case x < y :
            fmt.Println("`x` is smaller than `y` !")
        case x == y :
            fmt.Println("`x` is equal to `y` . ")
    }
    
    switch strings.ToLower(name) {
        case "rob" :
            fmt.Println("Your name is `rob`")
        case "andrew" :
            fmt.Println("Your name is `andrew`")
        case "erfan" :
            fmt.Printf("Your name is %s", name)
        default :
            fmt.Printf("Oh sorry your select isn't correct")
    }
}