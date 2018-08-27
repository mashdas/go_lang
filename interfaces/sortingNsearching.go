//Try sorting and searching too-->use silce,stringslice,intslice,custom Len,Swap,Less

package main

import (
	"fmt"
	"sort"
)

type douche struct{
	name string
	age int
}
func (d douche) String() string {
	return fmt.Sprintf("%v and I'm %d years old", d.name, d.age)
}


//making this collection implement the "Interface" interface,i.e making it implement the methods Len,Less,Swap
//and running sort.Sort on them
var telemarketers=[]douche{{"Jev",22},{"Malfoy",24},{"crabbe",27}}

type byAge []douche//---->To make byAge implement the "Interface" interface and convertong tele type to byAge type

func (x byAge) Len() int {return len(x)}
func (x byAge) Less(i,j int) bool{return x[i].age>x[j].age}
func (x byAge) Swap(i,j int) {x[i],x[j]=x[j],x[i]}
//We can use the above method to sort strings and ints too..but we have built in methods already
//It would have been like the following
//i:=[]int{1,2,4345,4,55,6}
//
//type stuff [] int
//func (s stuff) Len(s) int{return len(s)}
//func (s stuff) Swap(i,j int) {s[i],s[j]=s[j],s[i]}
//func (s stuff) Less(i,j int) bool{return s[i]>s[j]}
//sort.Sort(s(i))

func main(){
	douche_army:=[]douche{{"Twennifer",23},{"Ongo",42},{"Gablogian",43},
	                    {"Predre",33}}
	//USING JUST SLICE--->need to define the less method manually
	sort.Slice(douche_army,func (i,j int) bool{return douche_army[i].age>douche_army[j].age})//Use this instead of soo much code
	fmt.Println(douche_army)

	sort.Sort(byAge(telemarketers))
	//sort.Sort(sort.Reverse(byAge(telemarketers)))---->reverses the sorting order..probably changes the Less Function
	fmt.Println(telemarketers)

	//Searching a struct
	location:=sort.Search(len(telemarketers),func(i int) bool{return telemarketers[i].name=="Malfoy"})
	fmt.Println(location)
	//Sorting  a string
	gang:=[]string{"Frank","Charlie","mac","Dee","dennis"}
	sort.StringSlice(gang).Sort()
	fmt.Println(gang)
	vics:=[]string{"Cricket","Maureen","Artemis","McPoyle"}
	sort.Strings(vics)
	fmt.Println(vics)
	//Searching an int....Searching can only be done after sorting in an ascending order
	i:=[]int{1,2,4345,4,55,6}
	sort.IntSlice(i).Sort()
	loc:=sort.IntSlice(i).Search(4)
	fmt.Println(loc)


	}