# Ejercicio 1

Write a program that prints a number in binary, decimal, hexadecimal and octal format. Ask the user for the number (one single number) to print.

## Solution GO

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var number int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&number)
    fmt.Printf("Binary: %b\n", number)
    fmt.Printf("Decimal: %d\n", number)
    fmt.Printf("Hexadecimal: %x\n", number)
    fmt.Printf("Octal: %o\n", number)

// Other solution could be
    fmt.Printf(" %d\t %b\t %o\t %x\t", number, number, number, number)
}

```

