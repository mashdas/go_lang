package main

import "fmt"

func main(){
	k:=[][]int{}

	for j:=0;j<3;j++{
		m:=[]int{}
	for i:=0;i<4;i++{
		m=append(m,i)
	}
	k=append(k,m)
	}
	fmt.Println(k)

	//slice of slice of strings---->This time using make
	x:=make([][] string,3)//---->this will have created a empty slice of slice of strings of len 3==>[ [] [] [] ]===>#bla
	for i:=0;i<3;i++{
		y:=make([] string,5)//---->this will have created empty slice of strings---->[     ]==>5 blank spaces as specified
		for j:=0;j<5;j++{//if used append it would have been like this-->[     ,x,y,z]==>appendage after 5 blank spaces..same would apply in "bla"
			y[j]="c"
		}
		x[i]=y
	}
	fmt.Println(x)
	fmt.Println(x[0][1])

}