package main

import (
	"fmt"
)

func main() {
	arr1 := [5]int{1: 10, 2: 40}
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized
	arr4 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(len(arr4))
}
