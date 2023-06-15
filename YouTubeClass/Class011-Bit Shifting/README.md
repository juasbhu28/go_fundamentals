# Bit shiftign

## Bit shifting is a way to quickly multiply or divide by 2

Bit shifting in GO is done using the `<<` and `>>` operators.

It's useful to know that shifting left by 1 is equivalent to multiplying by 2, and shifting right by 1 is equivalent to dividing by 2.

## Bit shifting is fast

Bit shifting is a very fast way of multiplying or dividing by 2. It's faster than using multiplication or division, but of course it's also less flexible.

## Bit shifting is not the same as modulo

# Exercise

1. Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

2. Write a function which takes two integers and returns their sum as an integer. However, you must not use the + operator.

3. Write a function which takes two integers and returns their sum as an integer. However, you must not use the + or - operator. `´

# Solutions

## 1. Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

```go  
package main

import "fmt"

//The function half have two return values, the first is the half of the number and the second is a boolean value that indicates if the number is even or not.
func half(x int) (int, bool) {
    return x / 2, x%2 == 0
}

func main() {
    fmt.Println(half(1)) // 0, false
    fmt.Println(half(2)) // 1, true
}
```

## 2. Write a function which takes two integers and returns their sum as an integer. However, you must not use the + operator.

```go
package main

import "fmt"

func sum(x, y int) int {
    if y == 0 {
        return x
    }
    sum := x ^ y
    carry := (x & y) << 1
    return sum(sum, carry)
}

func main() {
    fmt.Println(sum(1, 1))
    fmt.Println(sum(2, 3))
}
```

## 3. Write a function which takes two integers and returns their sum as an integer. However, you must not use the + or - operator.

```go

package main

import "fmt"

func sum(x, y int) int {
    if y == 0 {
        return x
    }
    sum := x ^ y
    carry := (x & y) << 1
    return sum(sum, carry)
}

func main() {
    fmt.Println(sum(1, 1))
    fmt.Println(sum(2, 3))
}
```

