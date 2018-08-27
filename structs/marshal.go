package main

import (
	"encoding/json"
	"fmt"
)

type racoon struct{
	Name        string
	Age         int   `json:"wisdom-score""`//if you do small letter,i.e "age",it'll not be exported and you won't get that data processed
	Catchphrase string
	isRocket	bool  `json:"-""`//will ignore this field
}

func main(){


	d1:=racoon{"TRex",121,"noise",true}
	fmt.Println(d1)
	x,_:=json.Marshal(d1)
	fmt.Printf("%T",x)

	fmt.Println(x)
	fmt.Println(string(x))


	//TAKEAWAY=====>Marshal takes in a sturct and returns a slice of bytes(This is just the over the top view;for inner workings view docs)



}