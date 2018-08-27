package main

import "fmt"

func main(){
	x:=func(a []int,callback func(data int) bool) []int{
		xs:=[]int{}
		for _,y:=range  a{
			if callback(y){
				xs=append(xs,y)
			}

		}
		return xs
	}

	xs:=x([] int{1,2,3,4,4,5,6,6,7,8,8}, func(a int) bool{
		return a>4
	})
fmt.Println(xs)
}