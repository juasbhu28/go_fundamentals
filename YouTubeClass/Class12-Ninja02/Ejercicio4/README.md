# Exercice 4

Write a program that:
1. Assigns an int to a variable
2. Prints that int in decimal, binary and hex
3. Shifts the bits of that int over 1 position to the left, and assigns that to a variable
4. Prints that variable in decimal, binary and hex


## Solution GO

```go
package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
    b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
	