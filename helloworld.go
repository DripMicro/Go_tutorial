package main

import (
	"fmt"
)

func main() {
	var i = 15.5
	var txt = "Helo World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	var j = 15

	fmt.Printf("%b\n", j)
	fmt.Printf("%d\n", j)
	fmt.Printf("%+d\n", j)
	fmt.Printf("%o\n", j)
	fmt.Printf("%O\n", j)
	fmt.Printf("%x\n", j)
	fmt.Printf("%X\n", j)
	fmt.Printf("%#x\n", j)
	fmt.Printf("%4d\n", j)
	fmt.Printf("%-4d\n", j)
	fmt.Printf("%04d\n", j)
}
