package main

import "fmt"

func main() {
	x := func(a []int, callback func(data int)) {//callback is inplemented during the definition of boss function
		for _, y := range a {
			callback(y)
		}
	}

	x([]int{1, 2, 3, 4, 5, 6}, func(a int) {//callback is defined at  the time of calling the boss function
		fmt.Printf("%v\n", a*a*a)
	})
}

//=====================WITHOUT CALLBACK========================
//func callback(a int) {
//	fmt.Printf("%d\n",a*a*a)
//}
//
//func main() {
//	x := func(a []int) {
//		for _, y := range a {
//			callback(y)
//		}
//	}
//	x([]int{1,2,3,4,5,6})
//}