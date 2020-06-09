package main

import (
	"fmt"
	"strings"
)

//3
func repeatMe(words ...string) {
	fmt.Println(words)
}

//2
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//1
func multiply(a, b int) int {
	return a * b
}

func main() {
	// fmt.Println(multiply(2, 2))
	// totalLength, upperName := lenAndUpper("nico")
	// fmt.Println(totalLength, upperName)
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}
