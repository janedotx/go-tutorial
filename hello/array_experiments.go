package main

import "fmt"

func TestArray() {
	arr := []int{1, 2, 3}
	fmt.Println("this is the array before InnerTestArray")
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println((arr)[i])
	// }
	fmt.Printf("%v\n", arr)
	InnerTestArray((&arr))
	fmt.Println("this is the array after InnerTestArray and in TestArray")
}

func InnerTestArray(arr *[]int) {
	(*arr)[0] = 36
	fmt.Printf("%v\n", *arr)
	fmt.Println("this is the array after InnerTestArray")

	x := 5
	arr2 := make([]int, x)
	fmt.Printf("%v\n", arr2)
}