package main

import "fmt"

type vehicle struct{
	wheels int
	seats int
	doors int
	cost int
}

type tesla struct{
	vehicle
	engine string
	cost int
}

type duster struct{
	shitcar bool
	wheels int
	engine string
}
func main(){

	r1:=tesla{
		vehicle:vehicle{4,6,4,2000},
		engine:"electric",
		cost:4200,
		}
//or
	r3:=tesla{
		vehicle{4,6,4,2000}, "electric",
		4200,
	}
	print(r3)


r2:=duster{true,4,"Unicorn Tears"}
fmt.Println(r2)
fmt.Println(r1)


	}