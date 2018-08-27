package main

import "fmt"

type book struct{
	name string
	pages int
}

func (b book) showOff() string{
	return "My name is "+b.name+"I wrote This book!"
}

func sayStuff(a interface{}){//empty interface--->can pass in anything..but can't attempt to use any method on th concrete type
	fmt.Println(a)
}

func (b book) String() string{
	return fmt.Sprintf("%s has %d pages",b.name,b.pages)
}

func main(){

b1:=book{"ahgatha",234}
sayStuff(b1)
fmt.Println(b1)

}