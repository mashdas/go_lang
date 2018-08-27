package main

import (
	"fmt"
	"sort"
)

type p  [] string

//func (x p) Len() int {return len(x)}
//func (x p) Swap(i,j int) {x[i],x[j]=x[j],x[i]}
//func (x p) Less(i,j int) bool{return x[i]<x[j]}
func main(){

	var r=[] string {"Eenie","Meenie","miney","Xeon"}

	fmt.Println(r)
	q:=sort.StringSlice(r).Search("Xeon")
	fmt.Println(q)
	//
	//sort.Sort(p(x))
	//
	////sort.StringSlice(x).Sort()
	//fmt.Println(x)
	////
	//y:=sort.StringSlice(x).Search("miney")
	//fmt.Println(y)

}