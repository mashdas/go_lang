package main

import "fmt"

func main(){
	a:=0
	fmt.Scan(&a)
	k:=make([] int,a);
for i:=0;i<a;i++{
	fmt.Scan(&k[i])//if you declare the size then you can do this otherwise use append
}
fmt.Println(k)
}