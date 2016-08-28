package main

import (
    "fmt"
)

func  startWelcome(name string) (n string) {
    n = fmt.Sprintf("Welcome %s", name)
    return
}

func main() {
    fmt.Println(startWelcome("Erfan"))
}