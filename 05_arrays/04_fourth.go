package main

import (
    "fmt"
)

var (
    myArr [2][2]string
)

func main() {
   myArr[0][0] = "one"
   myArr[0][1] = "two"
   myArr[1][0] = "three"
   myArr[1][1] = "four"
   fmt.Printf("%+v\r\n", myArr)
   fmt.Printf("%#v", myArr)
}