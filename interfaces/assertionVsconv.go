package main

import "fmt"

func main() {
	var k interface{} = "Sydney"
	d:=66
	fmt.Println(k.(string) + "ombay")//Assertion---->Asserts that the concrete type that implements the empty interface is a string>Assertion is only for interfaces
	fmt.Println(string(d) + "ombay")//Conversion------>converts the decimal to a character
}