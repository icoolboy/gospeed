package main

import (
    "fmt"
)

type Master struct {
    Name    string
    Family  string
}

func main() {
    s := Master{Name : "Erfan", Family : "Akbarimanesh"}
    fmt.Println(s)
}