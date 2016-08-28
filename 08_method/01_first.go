package main

import (
    "fmt"
)

type Member struct {
    Name  string
}

func (m Member) SayHello() {
    fmt.Printf("Hi i'm %s", m.Name)
}

func main() {
    m := Member{"Mona"}
    m.SayHello()
}