package main

import (
    "fmt"
)

type Master struct {
    Name    string
    Family  string
}

func main() {
    s := Master{"Erfan", "Akbarimanesh"}
    fmt.Println(s)
}