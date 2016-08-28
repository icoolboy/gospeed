package main

import (
    "fmt"
    "strings"
)

type String string

func (s String) LowerCase() string {
    //return strings.ToLower(s) ---> this is fault, you must convert s to string
    return strings.ToLower(string(s))
}

func main() {
    fmt.Println(String("Hi I'm Erfan").LowerCase())
}