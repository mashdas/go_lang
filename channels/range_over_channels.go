package main

import "fmt"

func main(){
	c:=make(chan int)
	go func(){

		for i:=0;i<10;i++{
			c<-i
		}
		close(c)
	}()

	for x:=range c{
		fmt.Println(x)
	}
}

//Sends 0 prints 0
//Sends 1 prints 1
//...
//Sends 9 prints 9
//When nothing to send,the channel is closed and main is exited