# Numerics types in GO

## Integer vs Float64

In Go, there are two numeric types: integers and floating-point numbers. The predeclared integer types are `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, and `int64`. The predeclared floating-point types are `float32` and `float64`.

The `uint8` type, the set of all unsigned 8-bit integers (0 to 255) is also called `byte`, but `byte` is not a predefined name in Go. In case when you use this type, and assign a value greater than 255, the compiler will throw an error.


## Integer types

The `int` type is a signed integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, `int32`.

link: https://golang.org/ref/spec#Numeric_types


## Floating-point types

The floating-point types are `float32` and `float64`, corresponding to the IEEE-754 32-bit and 64-bit floating-point numbers.

link: https://golang.org/ref/spec#Floating-point_types

## Complex types

The predeclared complex types are `complex64` and `complex128`, whose underlying types are `float32` and `float64`, respectively.


## Numeric Conversions

A numeric value `x` can be converted to type `T` in any of these cases:

- `x` is representable by a value of type `T`.

- `x` is a floating-point constant, `T` is a floating-point type, and `x` is representable by a value of type `T` after rounding using IEEE 754 round-to-even rules. The constant `T(x)` is the rounded value of `x` (or, if `T`'s value set cannot represent `x`, the nearest representable value to `x`).

- `x` is an integer constant and `T` is a string type. The same rule as for non-constant `x` applies in this case.

```go

package main

import (
    "fmt"
    "math"
)

var x int

func main() {
    x = 3.14 // Error: constant 3.14 truncated to integer
    fmt.Println(x)

    //Convert float to int
    x = int(3.14)

    x = math.Pi // Error: constant 3.141592653589793 truncated to integer
    fmt.Println(x)
}

```

## Alias types

Exist two alias types: `byte` and `rune`. They are aliases for `uint8` and `int32` respectively. 

```go

package main

import (
    "fmt"
)

func main() {
    var x byte = 255
    fmt.Println(x) // 255

    var y rune = 2147483647
    fmt.Println(y) // 2147483647
}

```

## Print runtime type

```go

package main

import (
    "fmt"
    "runtime"
)

func main(){
    fmt.Println(runtime.GOOS) // linux
    fmt.Println(runtime.GOARCH) // amd64 -

   


}
x