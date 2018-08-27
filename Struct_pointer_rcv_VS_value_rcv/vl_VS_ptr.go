//value vs pointer receiver
package main

import "fmt"



type car struct{
	name string
	wheels int
	speed int
}

func (c car) intro() string{
	c.name="Honda"
	return c.name+"Rules!"
}

func (c *car) topSpeed(a int) int{
	c.speed=a
	return a
}

func main(){
	x:=car{"Pajero",4,100}
	fmt.Println(x.intro())//Value in x won't be changed
	fmt.Println(x)
	fmt.Println(x.topSpeed(200))//The value in x will be changed
	fmt.Println(x)
}

