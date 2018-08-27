package main

func stuff(x *int) int{
	*x=*x+21
	return *x
}


func main(){
	x:=20
	//var y *int=&x
	//fmt.Printf("%p",y)
	print(stuff(&x))


	}