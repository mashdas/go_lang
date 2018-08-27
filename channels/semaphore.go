package main

import (
	"fmt"
)


//many functions writing to the same channel
func main(){
c1:=make(chan int)
d:=make(chan bool)
	go func(){
		for i:=0;i<10;i++{
			c1<-i
			//time.Sleep(2*time.Millisecond)
		}
		d<-true
	}()

	go func(){
		for j:=0;j<10;j++{
			c1<-j
			//time.Sleep(4*time.Millisecond)
		}
		d<-true

	}()

	go func(){
		<-d
		<-d
		close(c1)
	}()

for x:=range c1{
	fmt.Println(x)
}
}