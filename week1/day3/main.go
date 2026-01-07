package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("ghj")
	math.Mod(33.3, 3)


/****************************************************************************/

var mp map[string]int

mp=make(map[string]int)

mp["abc"]=2
mp2:=map[string]int{
	"we":3,
	"www":34,
	"332":11,
 }

 fmt.Println(mp2)
 fmt.Println(mp)
arr1 := []int{3,1,2,4,5,6,3,1,2,4,6,8,9,2,123,34,64,33}
var mpToFindNonRepeatingEle map[int]int
mpToFindNonRepeatingEle =make(map[int]int)
 for v:= range arr1 {
	mpToFindNonRepeatingEle[arr1[v]]++;
 }
 for v:= range arr1 {
	
	if(mpToFindNonRepeatingEle[arr1[v]]==1){
	fmt.Println(arr1[v])
	break;}
 }
 

mp3:= map[string]int {
	"5":77,
}
mp3["c"] = 3              // write
fmt.Println(mp3["c"]) 

fmt.Println(mp3["5"])
fmt.Println(mp3["5"]) 
for i, v:= range mpToFindNonRepeatingEle { fmt.Println(i,v) 
	
	vi,okk :=mpToFindNonRepeatingEle[i]

	fmt.Println(vi,okk)
	
	} 
var mp6 map[string]int

fmt.Println(mp6["a"]) 










/***********************************************************************************************/





// maps done now functions 




}

func f1(num []int) {
	for i:=range num{
		num[i]*=2
	}
}

func f2(mp map[int]int){
	for k,v:=range mp{
		mp[k]=v+5
	}
}