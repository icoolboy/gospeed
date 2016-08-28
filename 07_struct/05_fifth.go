package main

import (
    "fmt"
)

type Master struct {
    Name      string
    Family    string
    Information
    IsEnable  bool
}

type Information struct {
    Age  uint8
    ExpertCategory string
}

func main() {
    s := Master{Family : "Akbarimanesh", Information: Information{20, "Go"}}
    fmt.Println(s)
    
    s.Name = "Erfan"
    fmt.Println(s.Name)
    fmt.Println(s.Family)
    
    s.IsEnable = true
    fmt.Println(s)
}