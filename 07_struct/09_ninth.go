package main

import (
    "fmt"
)

type User struct {
    Id            int
    Name, Family  string
    
    //Private fields
    profile
    isEnable      bool
}

type profile struct {
    Age   int
}

func main() {
    u := User{1, "Reza", "Alavi", profile{Age : 36}, false}
    fmt.Println(u)
}