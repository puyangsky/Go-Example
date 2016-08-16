package main

import "fmt"

func main() {
	fmt.Println("hello world")
	b := test(1)
	fmt.Println(b)
}

func test(a int) int {
	return a + 1
}
