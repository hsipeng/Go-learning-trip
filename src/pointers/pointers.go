package main

import "fmt"

func change(val *int) {
	*val = 55
}

func modify(sls []int) {
	sls[0] = 90
}
func main() {

	// 1. example
	// b := 255

	// var a *int
	// if a == nil {
	// 	fmt.Println("a is ", a)
	// 	a = &b
	// 	fmt.Println("a after initialization is", a)
	// }

	// 2. new(type)

	// size := new(int)
	// fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	// *size = 85
	// fmt.Println("new size value is ", *size)

	// 3. passing pointer to a function

	// a := 58

	// fmt.Println("value before change is ", a)

	// b := &a
	// change(b)
	// fmt.Println("value after change is ", a)

	// 4. arrary using slices

	a := [3]int{89, 90, 91}

	modify(a[:])
	fmt.Println(a)
}
