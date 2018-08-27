package main

import "fmt"

//var x=10
//func increment() int {
//	 x++
//	 return x
//}
//
//func main(){
//	print(increment())=------------->11
//	print(increment())--------------->12
//}

func main(){
	x:=10
	var stuff =func() int {
		var k=func() int{
			return 1
		}
		x++
		return x + k()
	}
	print(stuff())
	fmt.Printf("\n")
	print(stuff())
}

//For some reason you cant declare functions inside functions using a name and stuff..it has to be anonymous