package main

import (
	"fmt"
	"strconv"
	"reflect"
)

// Go does not have enums directly.
// We usually create enum-like constants with iota.
// type LogLevel int

const (
	LogError  = iota + 'a'
	LogWarning
	LogInfo
)

func main() {
	var level  = LogInfo

	fmt.Println("Error level:", string(LogError))
	fmt.Println("Warning level:", string(LogWarning))
	fmt.Println("Info level:", string(LogInfo))
	fmt.Println("Current level:", string(level))


	var num, err  = strconv.Atoi("123") // Convert string to int


	if(err != nil) {
		fmt.Println("Error converting string to int:", err)
		return
	}
	integ1 := strconv.Itoa(123) 
	
	
	fmt.Println(num)
	fmt.Println(integ1)

	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.TypeOf(integ1))	
	// Convert int to string



	
}
 