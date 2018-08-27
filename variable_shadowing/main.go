package main

import "fmt"

func a(x int) int{
	return x+10
}

func main(){
	//Assuming one is familiar with the variable declarations..let's proceed
	fmt.Printf("%d\n",a(90));
	a:=10;
	fmt.Printf("%d\n",a)
	//If I now call a(stuff)..i'll no longer be able to access the function cuz a now has been assigned to an int
}