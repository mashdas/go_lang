package main

import (
	"sync"
	"fmt"
	"time"
)
var mutex sync.Mutex
var Wg sync.WaitGroup
var Counter int

func main() {
	Wg.Add(2)
	go Incrementor("Foo:")
	go Incrementor("Bar")
	Wg.Wait()
	fmt.Println("Final Count:", Counter)

}

func Incrementor(s string){
	for i:=0;i<20;i++{
		time.Sleep(time.Duration(4*time.Millisecond))
		mutex.Lock()
		Counter++
		mutex.Unlock()
		fmt.Println(s,i,"Counter:",Counter)
	}
	Wg.Done()
}


