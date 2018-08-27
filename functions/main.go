package main

import "fmt"

func main(){
	x:="this is a string"
	stuff(x)
	a:=116
	fmt.Printf("%c",a)//will print out the character
}

func stuff(x string){
	for _,a:=range x{
		fmt.Printf("%v\n",a)//will print out the runes
	}



}