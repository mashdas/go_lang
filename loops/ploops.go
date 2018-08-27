package main

import "fmt"



func main(){
//	for i:=0;i<5;i++{
//		for j:=0;j<5;j++{
//			fmt.Printf("%d-%d\n",i,j)
//}
//
//	var x=func (){
//			print("UUUUUU")
//		}
//	x()
//}

	//while is replaced by for

	//k:=0;
	//for k<5{
	//	if k>5{
	//		break
	//	}
	//	print(k)
	//	k++
	//}
	//Printing only odd numbers
	//using "while" for

	x:=0

	for x<=50{
		x++
		if x%2==0{

			continue
		}
		fmt.Printf("%d\n",x)



		}


	}


