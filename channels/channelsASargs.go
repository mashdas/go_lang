package main

import "fmt"




func main(){
	d:=generator()
	fmt.Println("=============")
	k:=Summer(d)
	fmt.Println("=============")
	for x:=range k{
		fmt.Println(x)
	}
	fmt.Println("=============")



}

func generator() chan int {
	c := make(chan int)
	go func(){
		for x := 0; x < 10; x++ {
			c <- x
		}

		close(c)
	}()
	return c

}


func Summer(c chan int) chan int{
	sum:=0
	out:=make(chan int)
	go func(){
		for x:=range c{
			sum+=x
		}
		out<-sum
		close(out)
	}()


	return out
}