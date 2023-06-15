package main

import (
	"fmt"
)

//Go is a static code, means that the type of the variable is defined when the variable is created.
//The type of the variable can not be changed.
//The type of the variable can be inferred by the compiler.

var a int
//For create a date type we use the keyword type

type dinero int
var b dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 100
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b -> Error, because a is an int and b is a dinero
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
//Result:
// 42
// int
// 100
// main.dinero

