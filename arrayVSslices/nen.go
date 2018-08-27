package main

import "fmt"

func main(){
	mySlice:=[] string{"a","b","c","d"}
	fmt.Println(mySlice[2])
	fmt.Println(mySlice[1:2])
	fmt.Println(mySlice[0:2])
	fmt.Println("mySlice"[2])
	print("c")//--->c
	print('c')//--->some number(it's ascii value)
}