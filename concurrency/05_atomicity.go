package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

var Wgp sync.WaitGroup
var Ctr int64

func main() {
	Wgp.Add(2)
	go Inc("Foo:")
	go Inc("Bar")
	Wgp.Wait()
	fmt.Println("Final Count:", Ctr)

}

func Inc(s string){
	for i:=0;i<20;i++{
		time.Sleep(time.Duration(4*time.Millisecond))
		atomic.AddInt64(&Ctr,1)
		fmt.Println(s,i,"Counter:",Ctr)
	}
	Wgp.Done()
}


