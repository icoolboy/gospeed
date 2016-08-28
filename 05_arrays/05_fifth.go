package main

import (
    "fmt"
)

var (
    myArr [4]int
)

func main() {
   myArr = [4]int{10, 2, 3, 54}
   fmt.Println("Index[1] to Index[3] :", myArr[1:])
   
   fmt.Println("Index[3] to Index[3] :", myArr[3:])
   
   fmt.Println("Index[1] to Index[3] :", myArr[1:3])
}
