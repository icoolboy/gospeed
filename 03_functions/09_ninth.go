package main

import "fmt"

func myFunc(n, m int, callback func(int, int)) {
	callback(n, m)
}

func main() {
	myFunc(1, 2, func(n, m int) {
		fmt.Println(n + m)
	})
}
