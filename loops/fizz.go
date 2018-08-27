package main

import "fmt"

func main(){
	//for i:=0;i<50;i++{
	//	if i%3==0 && i% 5==0{
	//		fmt.Printf("%d--Fizz-Buzz\n",i)
	//	}
	//	if i%5==0 && i%3!=0{
	//		fmt.Printf("%d--Buzz\n",i)
	//	}
	//	if i%3==0 && i%5!=0{
	//		fmt.Printf("%d--Fizz\n",i)
	//	}
	//}
	k:=0
	fmt.Scanf("&k")
	print(k)
	for i:=0;i<50;i++{
		switch{
		case i%3==0 && i%5==0:
			fmt.Printf("%d-fizzbuzz\n",i)
		case i%5==0 && i%3!=0:
			fmt.Printf("%d--Buzz\n",i)
		case i%3==0 && i%5!=0:
			fmt.Printf("%d--Fizz\n",i)
		default:
			fmt.Printf("%d\n",i)
		}
	}
}