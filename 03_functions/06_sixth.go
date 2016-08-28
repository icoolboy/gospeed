package main

import (
    "fmt"
)

func  startWelcome(name, family string) (string, string) {
    return fmt.Sprintf("Name : %s", name), fmt.Sprintf("Family : %s", family)
}

func main() {
    n, f := startWelcome("Erfan", "Akbarimanesh")
    fmt.Println(n)
    fmt.Println(f)
}