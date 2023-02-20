package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Args)
	for _, url := range os.Args[1:] {
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			fmt.Println("bisaaa")
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			fmt.Println(" tidak bisaaa")
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
