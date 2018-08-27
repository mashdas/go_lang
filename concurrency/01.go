package main

import (
	"fmt"
	"sync"
)

func foo(){
	for i:=0;i<45;i++{
		fmt.Println("FOO:",i)
	}
	x.Done()
}
func bar(){
	for i:=0;i<45;i++{
		fmt.Println("BAR",i)
	}
	x.Done()

}

var x sync.WaitGroup
func main(){
	x.Add(2)
	go foo()
	go bar()
	x.Wait()
	fmt.Println("Done")
}

//What happens is if i don't use the wait group,main will be executed,as the function call will be moved to stack,the
//main func exits-->nothing will get done
//=========================================
//In this case,after the use of weight group,After the end of every function,delta in the add function is decremented
//when it's zero,main will exit after running the rest of the code