package main

import (
    "fmt"
)

type Master struct {
    Name      string
    Family    string
    IsEnable  bool
}

func main() {
    s := Master{Family : "Sarabi"}
    fmt.Println(s)
    
    s.Name = "soroush"
    fmt.Println(s.Name)
    fmt.Println(s.Family)
    
    s.IsEnable = true
    fmt.Println(s)
}