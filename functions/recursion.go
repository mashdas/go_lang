package main

import "fmt"

func fibo(a int){
	x,y:=0,1
	for i:=0;i<a;i++{
		x,y=y,x+y
		fmt.Println(x,"/n")

	}

}


func recur_fibo(a int) int{
	if a==0 || a==1{
		return 1
	}else{
		return recur_fibo(a-1)+recur_fibo(a-2)
	}
}

func main(){
	fibo(10)
	fmt.Println(recur_fibo(10))

}