package main

func main(){
	x:=10
	var increment=func () func() int{
		return func() int{
			x++
			return x
		}
	}
	ment:=increment()
	print(ment())
	print(ment())
}