package main

import (
	"fmt"

	"github.com/rkh/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	dictionary.Search(baseWord)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		word, err2 := dictionary.Search(baseWord)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(word)
		}
	}
}
