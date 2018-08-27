package main

import "fmt"

func main(){

	n:=10
	c:=make(chan int)
	d:=make(chan bool)

	go func(){
		for i:=0;i<n;i++{
			c<-i
			}
		close(c)
	}()

	for i:=0;i<n;i++{
		go func(){
			for n:=range c{
				fmt.Println(n)
			}
			d<-true
		}()

	}

	for i:=0;i<n;i++{
		<-d
	}


	}

