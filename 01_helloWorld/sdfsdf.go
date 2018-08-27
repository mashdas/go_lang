package main

import (
	"fmt"
	"sort"
)

func main() {



	d:=[]string{"Huckleberry","Tom","Gwendolyn","Msa","Onyx"}
	x:=sort.StringSlice(d).Search("Onyx")//returns 4 this works though
	y:=sort.StringSlice(d).Search("Tom")
	z:=sort.StringSlice(d).Search("Gwendolyn")
	a:=sort.StringSlice(d).Search("Msa")
	r:=sort.StringSlice(d).Search("Huckleberry")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(r)

	fmt.Println(d)
}