package main

import "fmt"

func main() {
	k := map[int]string{}
	k[1], k[2], k[3], k[4], k[5] = "Hydrogen", "Helium", "lithium", "Berylium", "Boron"

	for key,value:=range k{
		fmt.Println("Key-",key," Value-",value)
	}

	}