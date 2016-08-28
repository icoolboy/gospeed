package main

import (
    "fmt"
)

var (
    x, y int = 1, 2
    appName, version, productId = "MyFirstGoApp", "1.0.0", 118072
)

func main() {
    name, family := "Erfan", "Akbarimanesh"
    fmt.Printf("%d %d\r\n", x, y)
    fmt.Println(name, family)
    fmt.Printf("\r\nAppName : %s \r\nVersion : %s \r\nProduct Id : %d", appName, version, productId)
}
