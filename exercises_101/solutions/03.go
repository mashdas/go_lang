package solutions

import "fmt"

func Enquire(){
	i:=""
	print("What's yo name")
	fmt.Scan(&i)
	fmt.Printf("Hello Mr.%s",i)
}

func Showoff(){
	print("Enter two numbers")
	var i,j int
	fmt.Scan(&i,&j)
	switch{
	case i>j:
		fmt.Printf("%d",i%j)
	case j>i:
		fmt.Printf("%d",j%i)
	}
}

func Evens(){
	for i:=0;i<=100;i++{
		if i%2==0{
			fmt.Printf("%d-EVEN\n",i)
		}else{
			continue
		}
	}
}

func Sum41(){
i:=0
for j:=0;j<1000;j++{
	if j%3==0 || j%5==0{
		i+=j
	}

}
	print(i)
}