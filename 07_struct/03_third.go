package main

import (
    "fmt"
)

type Master struct {
    Name    string
    Family  string
}

func main() {
    s := Master{}
    fmt.Println(s)
    
    s.Name = "Erfan"
    s.Family = "Akbarimanesh"
    fmt.Println(s.Name)
    fmt.Println(s.Family)
    fmt.Println(s)
}