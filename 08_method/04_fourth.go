package main

import (
    "fmt"
)

type Member struct {
    Name   string
}

//get arg from input of func
func (m Member) sayHello(family string) string {
    return fmt.Sprintf("Hi i'm %s %s", m.Name, family)
}

func main() {
    m := Member{Name : "Amir"}
    fmt.Println(m.sayHello("Amini"))
}