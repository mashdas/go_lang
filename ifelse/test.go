package main

func main(){


	if food:="maggi" ;len(food)==5{//This is done to keep the scope tight..the variable food now is only accessible
		print("man's not hot")//inside the if block
		print(food)
	}

	//print(food)=====>Food is not accessible here
	//The above is not equivalent to:
	x:="maggi"
	if len(x)==5{
		print("The two scenarios are diff.")
	}
	print(x)//x is.

}