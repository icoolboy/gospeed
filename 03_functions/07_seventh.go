package main

import (
    "fmt"
)

func main() {
    welcome := func(name string) string {
        return "Welcome " + name
    }
    
    fmt.Println(welcome("Erfan"))
}
