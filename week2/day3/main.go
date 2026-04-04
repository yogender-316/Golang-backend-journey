package main

import (
	"fmt"
)

func main() {
	fmt.Println("Today we gonna learn switch and maps")

	// Define a map of products and their base prices
	// today we gonna use types as well
	 //pro := 3

	pro:="3"
	// not allowed to reassign a variable with a different type
fmt.Printf("pro: %v, type: %T\n", pro, pro)
// but we the help of interface{} we can achieve that
	var intExample interface{} = 3
	var stringExample interface{} = "3"
	fmt.Printf("intExample: %v, type: %T\n", intExample, intExample)
	fmt.Printf("stringExample: %v, type: %T\n", stringExample, stringExample)

	/// now we gonna change its type
	intExample = "Now I'm a string"
	stringExample = 100
	fmt.Printf("intExample: %v, type: %T\n", intExample, intExample)
	fmt.Printf("stringExample: %v, type: %T\n", stringExample, stringExample)

	// switch on types
	switch v := intExample.(type) {
	case int:
		fmt.Println("intExample is an int:", v)
	case string:
		fmt.Println("intExample is a string:", v)
	default:
		fmt.Println("intExample is of unknown type")
	}	

	products := map[string]float64{
		"ProductA": 10.0,
		"ProductB": 20.0,
		"ProductC": 30.0,
	}
	switch v := products["ProductA"].(type) {
	case float64:
		fmt.Println("Price of ProductA is a float64:", v)
	default:
		fmt.Println("Price of ProductA is of unknown type")
	}

	// how to correct above error
	price, exists := products["ProductA"]	
	if exists {
		fmt.Printf("Price of ProductA: $%.2f\n", price)




	}

	var pro2 interface{} = 3
	switch v := pro2.(type) {
	case int:
		fmt.Println("pro2 is an int:", v)
	case string:
		fmt.Println("pro2 is a string:", v)
	default:
		fmt.Println("pro2 is of unknown type")
	}	

	// in interviews i'll not be getting help to correct my code so i need to be careful while writing code and make sure that it is correct and compiles without errors

	// tips to be more careful while writing code:

	// 1. Always check for errors and handle them properly
	// 2. Use proper variable names and avoid using single letter variable names
	// 3. Use comments to explain your code and make it more readable
	// 4. Test your code thoroughly before submitting it
	// 5. Use a linter to check for common mistakes and improve code quality

	// wait tell me for them if i say im a go developer then they will not expect me to write code without vs code right
	// but if they ask me to write code without vs code then i should be able to do that as well because in interviews they might ask me to write code on a whiteboard or in a text editor without any syntax highlighting or code completion features	
	


}