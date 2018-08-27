package main

import (
	"strconv"
	"fmt"
)

type man struct{
	muggle bool
	weapon string
	warcry string
}

type wizard struct{
	man
	spell string
}

func (w wizard) attack()string{
	return w.spell
}

func (m man) attack(a string) string{
	return m.warcry+"I've got "+a+" "+m.weapon

}

func main() {
	//
	//Harry := wizard{man{false, "wand", "AAARRRRHHHH"}, "expelliarmus"}
	//fmt.Println(Harry.spell)
	//fmt.Println(Harry.man.attack("two"))

	k:=78
	fmt.Printf("%c",k)//Prints the letter or rune
	fmt.Println(strconv.Itoa(k))//prints "78"

	//let's see dem pointers
	Vernon:=&man{true,"tea-spoon","No posts on SUnday"}
	fmt.Println(Vernon)
	fmt.Println(*Vernon)
	fmt.Println(Vernon.weapon)
	fmt.Println(Vernon.attack("one"))
}