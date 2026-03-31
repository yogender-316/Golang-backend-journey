package main

import (
	"fmt"
	
)


func main() {

	fmt.Println("Today we gonna learn switch and maps")
	// Switch statement in Go
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	default:
		fmt.Println("Midweek")
	}



	// we gonna switch on types as well
	var x interface{} = 42
	switch v := x.(type) {
	case int:
		fmt.Println("x is an int:", v)
	case string:
		fmt.Println("x is a string:", v)
	default:
		fmt.Println("x is of unknown type")
	}


	// this is the one way`
	// another way to do type switch
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of unknown type")
	}
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("i is an int:", v)
		case string:
			fmt.Println("i is a string:", v)
		default:
			fmt.Println("i is of unknown type")
		}
	}
	checkType(100)
	checkType("Hello")
	checkType(3.14)


	// Now we will learn how to know type of variable in go
	var a int = 10
	var b string = "Hello"
	// var c float64 = 3.14
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", checkType)

	// how to know the type of a variable in an if statement
	if _, qk := interface{}(a).(int); !qk {
		fmt.Println("a is an int")
	}	 else {
		fmt.Println("a is not an int")
	}	
a=1
		var bc float64 = 1.14
		if a==int(bc) {
			fmt.Println("a and b are equal")
		} else {
			fmt.Println("a and b are not equal")
		}
  

		randomValue := interface{}(42.0)

		switch v := randomValue.(type) {
		case int:
			fmt.Println("randomValue is an int:", v)
		case float64:
			fmt.Println("randomValue is a float64:", v)
		default:
			fmt.Println("randomValue is of unknown type")
		}
		randomValue = "Hello, Go!"

		// we gonna check type assertions in if statements
		if val, ok := randomValue.(int); ok {
			fmt.Println("randomValue is a int:", val)
		} else if val, ok := randomValue.(float32); ok {
			fmt.Println("randomValue is a float32:", val)
		} else if val, ok := randomValue.(string); ok {
			fmt.Println("randomValue is a string:", val)
		}	else {
			fmt.Println("randomValue is of unknown type")
		}
		randomValue = 3.14

		// we gonna check type assertions in switch statements	


		switch v := randomValue.(type) {
		case int:
			fmt.Println("randomValue is an int:", v)
		case float64:
			fmt.Println("randomValue is a float64:", v)
		case string:
			fmt.Println("randomValue is a string:", v)
		default:
			fmt.Println("randomValue is of unknown type")
		}

		fmt.Printf("Type of randomValue: %T\n", randomValue)
		fmt.Println(randomValue)
		
		randomValue = func() { fmt.Println("Hello from a function!") }

		switch v := randomValue.(type) {
		case int:
			fmt.Println("randomValue is an int:", v)
		case float64:
			fmt.Println("randomValue is a float64:", v)
		case string:
			fmt.Println("randomValue is a string:", v)
		case func():
			fmt.Println("randomValue is a function")
			v()
		default:
			fmt.Println("randomValue is of unknown type")
		}	
		fmt.Println(randomValue)

		fmt.Printf("Type of randomValue: %T\n", randomValue)

}