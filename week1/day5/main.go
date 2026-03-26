package main
import "fmt"
func main () {

	
	var a rune  =12344

	for i:=0; i<10; i++{
		fmt.Printf("%c\n",a)
		a++
	}



	// slice 
	s := []int{1,2,3}


	for i:=0;i<5;i++ {
		fmt.Println(s[i])
	}


	// struct syntax 
	type Person struct {
		abc int
		str string
	}

	p := Person{
    str: "Vishnu",
    abc: 21,
	}

	fmt.Println(p)
}

func calc(a,b int) (int,int){
    return a+b, a-b
}