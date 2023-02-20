package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	a := 0
	counts := "ABb"
	input := bufio.NewScanner(strings.NewReader(counts))
	for input.Scan() {
		a++
		fmt.Println(a)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
