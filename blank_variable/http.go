package main

import (
	"net/http"
	"io/ioutil"
	//"fmt"
	"fmt"
)

/*func main(){
	res,err:=http.Get("https://www.golang-book.com/books/intro/1#section1")
	if err!=nil{
		log.Fatal(err)
	}

	page,err:=ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s",page)
	}*/

//blank variable is used mainly when we have to store something but we don't wanna use it.For instance in the above scenarion,if we don't want to deal with
//the error,we can use the blank identifier ,using any other identifier would cause an error---->Example follows
func main(){
	res,_:=http.Get("https://www.golang-book.com/books/intro/1#section1")
	page,_:=ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s",page)
}