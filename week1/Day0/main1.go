package main

import "fmt"

func main() {

	var a rune = 12344

	for i := 0; i < 10; i++ {
		fmt.Printf("%c\n", a)
		a++
	}
}
