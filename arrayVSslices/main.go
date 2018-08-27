package main

import "fmt"

func main(){

	var k [10]int
	for i,_:=range k{
		k[i]=i+21
	}


	m:=[] int {1,2,3,4,5}
	for i,_:=range m{
		m[i]=i+21
	}

	fmt.Printf("%v",k)
	fmt.Printf("%v",m)
}