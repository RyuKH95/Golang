package main

import "fmt"

func main() {
	a := 2
	b := &a
	c := &b
	a = 5
	fmt.Println(a, *b, **c)
	*b = 10
	fmt.Println(a, *b, **c)
}
