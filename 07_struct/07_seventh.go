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
    s := Master{Family : "Akbarimanesh", Name : "Erfan"}
    fmt.Println(s)
    
    s.Name = "Erfan"
    fmt.Println(s.Name)
    fmt.Println(s.Family)
    
    s.Information.Age = 20
    s.Information.ExpertCategory = "Go"
    s.IsEnable = true
    fmt.Println(s)
}