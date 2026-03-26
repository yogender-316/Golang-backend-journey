package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	j:=0
	for j < 5 {
		fmt.Printf("This will run indefinitely\n")
		j++
		 if j == 0 {
			fmt.Printf("Breaking out of the loop\n")
			break
		}
	}

	for x:=0; x < 5; {
		x++
		fmt.Printf("This is the %d iteration\n", x)
	}
}

