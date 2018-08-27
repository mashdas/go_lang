package main

import "fmt"

type Human struct{
	name string
	age int
	gender string

}

type social interface{
	sayHi() string

}

func (x Human) sayHi() string{
	return "Hi i'm"+x.name
}

func (x Human) Showoff(a string) string{
	return "My name is"+x.name+"I am a colossal"+a
}

func greet (z social){
	fmt.Println(z.sayHi)//Cannot use methods that are not used in the interface
	fmt.Println()
	//z.Showoff("Orewa")//---->will give an error

}

func main(){

p1:=Human{"Abed",24,"Male"}
fmt.Println(p1.Showoff("Douche"))
fmt.Printf("%T\n",p1)
var k social
fmt.Println(k)
k=p1
fmt.Printf("%T\n",k)
fmt.Printf("%v\n",k)




}

