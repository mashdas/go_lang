package main

import (
	"fmt"
	"sort"
)

//type p  [] string
//
//func (x p) Len() int {return len(x)}
//func (x p) Swap(i,j int) {x[i],x[j]=x[j],x[i]}
//func (x p) Less(i,j int) bool{return x[i]<x[j]}
func main(){

	x:=[] string {"Eenie","Meenie","miney","Xeon"}
	z:=sort.StringSlice(x).Search("Eenie")
	fmt.Println(z)
	//sort.StringSlice(x).Sort()
	//fmt.Println(x)
	//
	//y=sort.StringSlice(x).Search("Zam")
	//fmt.Println(y)

}