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
    m2 := new(Member)
    m2.Family = "Molavi"
    fmt.Println(m.SayHello())
    m.Family = "Rahmani"
    fmt.Println("My Family : ", m.whatIsMyFamily())
    
    family := m2.whatIsMyFamily()
    fmt.Println(family)
}