package main

import (
	"sort"
	"fmt"
)

func main() {

	r := []string{"Eenie", "Meenie", "miney", "Mo"}
	sort.StringSlice(r).Sort()
	fmt.Println(r)
	x:=sort.SearchStrings(r,"Eenie")
	fmt.Println(x)
//////////////////////////
//	r := []string{"Eenie", "Meenie", "miney", "Mo"}
//	sort.StringSlice(r).Sort()
//	fmt.Println(r)
//	x:=sort.StringSlice(r).Search("miney")
//	fmt.Println(x)

//to search as it is

k:=sort.Search(len(r),func(i int) bool{ return r[i]=="Mo"})
fmt.Println(k)

}