package main

import (
	"fmt"
	"math"
)
type circle struct{
	radius float32
}

//func (c *circle) area() float32{
//	return math.Pi*c.radius*c.radius---------------------------->BLOCK X
//}
func (c circle) area() float32{
	return math.Pi*c.radius*c.radius
}

type shape interface{
	area() float32
}

func info(z shape){
	fmt.Println("area:",z.area())
}

func main(){
c:=circle{5}
//INCASE OF VALUE RECEIVERS
info(c)//WORKS
info(&c)//WORKS

//INCASE OF POINTER RECEIVERS
//info(c)//DOESN"T WORK-------------------------------------------->CORRESPONDING X
//info(&c)//WORKS

//WHY--->BECAUSE OF HOW GO HANDLES UNTYPED CONSTANTS
}