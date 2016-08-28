package main

import (
    "fmt"
)

var (
    myArr []float64
    secondArr []string    
)

func main() {
    arr := []bool{true, false, true, true, true, false, false}
    fmt.Println(arr[4])
    fmt.Printf("Len of myArr : %d, secondArr : %d, arr : %d", len(myArr), len(secondArr), len(arr))
    fmt.Println(myArr)
    fmt.Println(secondArr)
}