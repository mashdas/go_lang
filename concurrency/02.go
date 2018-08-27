package main

import (
	"fmt"
	"sync"
	"time"
)

func Foo(){
	for i:=0;i<45;i++{
		fmt.Println("FOO:",i)
		time.Sleep(time.Duration(3*time.Microsecond))//For highlighting Concurrency
	}
	x.Done()
}
func Bar(){
	for i:=0;i<45;i++{
		fmt.Println("BAR",i)
	}
	x.Done()

}

var x sync.WaitGroup
func main(){
	x.Add(2)
	go Foo()
	go Bar()
	x.Wait()
	fmt.Println("Done")
}
