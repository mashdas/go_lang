package main

import "fmt"

//====================VariadicFunctions===============
func main(){
	n:=avg(43,56,87,12,45,57)
	fmt.Println(n)
}

func avg(data...float64) float64{//more like tuple packing
	fmt.Printf("%v\n",data)
	var total float64
	for _,k:=range data{
		total+=k
	}
	return total/float64(len(data))
}
//===============Variadic Parameters========================
//func main(){
//	a:=[]float64{21,23,34,56,78,89}
//	n:=dstuff(a...)//like tuple unpacking in python
//	fmt.Printf("%f",n)
//}
//
//func dstuff(data ...float64) float64{
//var total float64
//for _,v:=range data{
//	total+=v
//	}
//return float64(total)/float64(len(data))
//}

//=================================================

