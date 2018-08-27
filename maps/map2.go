package main

import "fmt"

func main(){

var m map[string]int

x := m

//m["hello"] = 42//will shit the bed...map is nil this assignment is a statement,panic is the only recourse

fmt.Println(x["hello"]==0)//True..try to assign something,will shit the bed

//deleting stuff
dict:=map[int] string{}
dict[0],dict[1],dict[2],dict[3]="Map","is","like","dict"

delete(dict,2)
if value,stat:=dict[2];stat{
	fmt.Println("Exists:",stat)
	fmt.Println("Value:",value)
}else{
	fmt.Println("Sorry it doesn't exist")
}

}
