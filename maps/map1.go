package main

import "fmt"

func main(){
	k:=make(map[string] int)
	k["a"]=1
	k["b"]=2
	k["c"]=3
	fmt.Println(k)

	//Without using  MAKE
	//WAY 1
	x:=map[string] int{"X":1,"Y":2}
	fmt.Println(x)
	//WAY 2
	var z map[string] int//nil map can't assign values like key-value pairs==>needs something like append for slice
	print(z)
	//DELETING STUFF
	delete(k,"a")//will delete the key 'a' and it's corresponding  value from the map 'k'
	fmt.Println(k)
	//Checking for the deleted key
	_,stat:=k["a"]//this returns the corresponding value and a boolean depending on the presence/absence of the said key
	print(stat)
	//CHECK
	a,b:=k["b"]
	print(a,b)//returns the value and a boolean
	//WAY 3
	r:=map[string] int{}
	r["aard"]=1
	r["vark"]=2
	fmt.Println(r)





}