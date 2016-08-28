package main

import "fmt"

func welcome(name string) {
	fmt.Println(name)
}

//without defer
func main() {
	welcome("Erfan")
    welcome("Mohsen")
}
