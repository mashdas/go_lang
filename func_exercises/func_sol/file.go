package func_sol

func Half(a int) (int,bool){
	return a/2,a%2==0

}

func Stuff(x int){
	a,b:=Half(x)
	print(a,b)
}


var M=func (x int){
	a,b:=Half(x)
	print(a,b)
}

func Greatest(data...int) int{
	greatest:=0
	for _,x:=range data{
		if x>greatest{
			greatest=x
		}
	}
	return greatest
	}


func Foo(data...int) int{
	total:=0
	for _,x:=range data{
		total+=x
	}
	return total
}