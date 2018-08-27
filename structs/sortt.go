package main

import (
	"sort"
	"fmt"
	"strconv"
	"strings"
)


//sorting structs

type doge struct{
	name string
	age int
}

func (d doge) woof() string{
	return "Woof,I'm "+d.name+" and I'm "+strconv.Itoa(d.age)+" years old"
}

type byAge []doge

func (x byAge) Len() int{return len(x)}

func (x byAge) Less(i,j int) bool {return x[j].age>x[i].age}

func (x byAge) Swap(i,j int) {x[i],x[j]=x[j],x[i]}

func main(){
	d:=[]string{"Huckleberry","Tom","Gwendolyn","Msatubu","Onyx"}
	fmt.Println(d)
	//sort.Sort(sort.StringSlice(d))
	fmt.Println(d)
	x:=sort.StringSlice(d).Search("Huckleberry")//returns the index of the search word after asc sort--->Here StringSlice implements the methods of 'Interface' interface in this collection
	fmt.Println(x)//Here x is not sorted and as such the value returned will be stupid
	sort.Sort(sort.StringSlice(d))
	x=sort.StringSlice(d).Search("Huckleberry")
	fmt.Println(d)
	fmt.Println(x)

	//sort.StringSlice(d).Sort()
	//fmt.Println(d)
	//x=sort.StringSlice(d).Search("Onyx")
	//fmt.Println(x)
	////Descending order
	//sort.Sort(sort.Reverse(sort.StringSlice(d)))
	//fmt.Println(d)

	//sort.Strings(d)//---->Sorts a slice of strings
	//fmt.Println(d)


	//SORTING STRUCTS---->WAY 1
	puppers:=[]doge{
		{"Shiba Inu",2},
		{"Husky",3},{"corgi",1},{"labrador",1},
	}
	sort.Sort(byAge(puppers))
	fmt.Println(puppers)

	//SORTING STRUCTS---->WAY 2(derivative)
	col:=byAge{{"akiba",2},{"rinu",1},{"ashi",3}}
	sort.Sort(col)//sort.Sort() expects something that implements the "Interface" interface and col is a collection that does implement the Interface interface via byAge
 	print("==========")
	fmt.Println(col)


	//there's an another way

	sort.Slice(puppers,func (i,j int) bool{return strings.ToUpper(puppers[i].name)<strings.ToUpper(puppers[j].name)})
	fmt.Println(puppers)

	//sorting a slice of int
	sliofint:=[]int{10,21,22,33,990,321}
	sort.Sort(sort.Reverse(sort.IntSlice(sliofint)))//Descending Order
	fmt.Println(sliofint)
	sort.Sort(sort.IntSlice(sliofint))//Ascending order
	fmt.Println(sliofint)




}

//NOTES:For using the sort package to sort there are two ways---->1>sort.Sort   2>sort.Slice
//For using sort.Sort one has to implement the "Interface" interface.To do that one has to incorporate 3 methods on a collection.Len(),Less(i,j int) bool,Swap(i,j int)
//when those methods are declared on the collection(x),a collection of the to be-sorted items is converted to type x.And they are sorted

//sort.Slice--->I'd rather use this one.requires no conversion,can be used directly,will have to declare and anonymous Less function..on the basis of which the collection will
//be sorted
