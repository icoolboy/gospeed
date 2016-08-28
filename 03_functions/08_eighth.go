package main

import "fmt"

func myFunc(n int, callback func(int)) {
	callback(n)
}

func main() {
	myFunc(1, func(n int) {
		fmt.Println(n)
	})
}
