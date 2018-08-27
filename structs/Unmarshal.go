package main

import (
	"fmt"
	"encoding/json"
	)

type Person struct{//One simolly can't loop over struct
	Name string
	Age  int
	Gender string
}

func main(){


	var L Person
	fmt.Println(L.Name)

	data:=[]byte(`{"Name":"Logic", "Age":32, "Gender":"m"}`)
	json.Unmarshal(data, &L)
	fmt.Println(L.Name)

	//k:=[]byte("Tjissis isisdfsdf sdf sdf sdf")
	//fmt.Println(k)

}
