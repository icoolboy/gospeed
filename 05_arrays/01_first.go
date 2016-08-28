package main

import (
    "fmt"
)

func main() {
   var myArr [10]int
   fmt.Println(myArr)
   myArr[0] = 12
   myArr[9] = 20
   fmt.Println(myArr[0], myArr[9])
   fmt.Printf("Array : %v", myArr)
}