package  main



func main(){
	x:=func (){
		print("hello")
	}
	x()

	y:=func (a string) func(b string) string{
		return func(b string) string{
			return a+b
		}

	}

	z:=y("Hello")
	print(z("Manish"))

}