package main

import "fmt"

func main(){
	const(
	_=iota
	a=1 << (iota*10)
	b=1 << (iota*10)
	)

	fmt.Printf("%d\n",a)
	fmt.Printf("%d",b)

	const k=10
	print(k)
	const p="20"
	print(p)
}