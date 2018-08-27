package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){

	url:="https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html"
	res,_:=http.Get(url)
	defer res.Body.Close()
	page,_:=ioutil.ReadAll(res.Body)
	fmt.Printf("%s",page)


}








