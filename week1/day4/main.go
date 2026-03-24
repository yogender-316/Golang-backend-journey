package main

import (
	"fmt"
)
 


type user struct {
	ID int
	Name string
	Email string
	Age int
}
func main() {
	fmt.Println("Hello world this is day4")

	u:=user{
		ID: 34,
		Name: "qwerty",
		Email: "qwerty111",
		Age: 23,
	}
	fmt.Println(u)
	fmt.Println(u.Name)

}