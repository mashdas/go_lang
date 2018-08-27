package main

import "fmt"

func main(){
	mySlice:=make([]int,0,5)//equlvalent to new([5]int)[0]

	fmt.Println("==========================")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Printf("Capacity-%d",cap(mySlice))
	fmt.Println("===========================")

	for i:=0;i<20;i++{
		mySlice=append(mySlice,i)
		fmt.Println("Length:" ,len(mySlice), "Capacity:",cap(mySlice))
	}


}