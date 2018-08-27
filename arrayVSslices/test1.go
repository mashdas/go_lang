package main

import "fmt"

func main(){
	//Trying to take input and shove it inside a slice
	print("Enter the size of the array")
	var a int
	fmt.Scan(&a)
	k:= []int{}//since this has length and capacity 0;elements can only be added via append
	for i:=0;i<a;i++{
		m:=0
		fmt.Scan(&m)
		k=append(k,m)//you have to use append when you want to add an element to the slice not k[some_index]=some_value
	}
	fmt.Println(k)


}

//TAKEAWAY=========>OUTSIDE THE LENGTH OF  THE SLICE;ELEMENTS CAN ONLY BE ADDED BY USE OF APPEND