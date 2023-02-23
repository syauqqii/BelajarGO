package main

import "fmt"

var a = 10

func s() int {
	c := 12
	return c
}

func main() {
	b := 11
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s())
}
