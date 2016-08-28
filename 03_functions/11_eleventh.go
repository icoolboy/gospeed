package main

import "fmt"

func welcome(name string) {
	fmt.Println(name)
}

//with defer
func main() {
	defer welcome("Erfan") //Changing the order of execution
    welcome("Mohsen")
}
