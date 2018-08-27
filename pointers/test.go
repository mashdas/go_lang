package main

func x(a *int){
	*a=10
}

func y(a int){
	a=0
}

func main(){
	var m int
	q:=20
	x(&q)
	print(q)

	y(m)
	print(m)


}