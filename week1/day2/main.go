package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("dfghk")
	var arg1 string ="string"
	if ah := 23;  to_string(ah)==arg1{
		fmt.Print("qwerty")
	}	else {
		fmt.Println(ah)
	}

	i :=3
	for it:= 0 ; it< 5; i++{
		fmt.Println(i)
		if i==34 {
			break;
		}
	}
	
	//making array 
	//  var name declaration then array symbol then data type
	arr := []int{20,23,21}
	fmt.Println(arr)






	// now functions over slices/array


	number1 := [] float32 { 34 ,342.5, 44,556, 661}

	for _,v:= range number1 {
		fmt.Println(v)
	}
	var count int
	count = -1
	count++
	for i:=0;i<len(number1);i++{
		fmt.Println("index: ", i, "Value : ", number1[i])

		if math.Mod(float64(number1[i]),2)==0{
			fmt.Println(math.Mod(float64(number1[i]),2))
			count++
		}
		fmt.Println(math.Mod(float64(number1[i]),2))
		fmt.Println("Count at ",i," index: ",count)
	}
	// for i:=
}

func  to_string( v interface  {}) string {
	return fmt.Sprint(v)
}
