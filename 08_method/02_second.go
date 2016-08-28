package main

import (
    "fmt"
)

type Member struct {
    Name, Family string
}

func (m Member) SayHello() string {
    return fmt.Sprintf("Hi i'm %s", m.Name)
}

func (m Member) whatIsMyFamily() string {
    if m.Family == "" {
        return "Oh i haven't family"
    }
    return m.Family
}

func main() {
    m := Member{Name : "Mona"}
    fmt.Println(m.SayHello())
    m.Family = "Pahlavan"
    fmt.Println("My Family : ", m.whatIsMyFamily())
}