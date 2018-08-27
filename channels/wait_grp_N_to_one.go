package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup


func main(){
	c:=make(chan int)
	wg.Add(2)

	go func(){
		for i:=0;i<10;i++{
			c<-i
		}
		wg.Done()
	}()

	go func(){
		for j:=0;j<10;j++{
			c<-j
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for x:=range c{
		fmt.Println(x)
	}
}