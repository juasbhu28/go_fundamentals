# Fudamentals Strings types in GO

Strings in GO are a sequence of bytes (uint8) and are immutable.

String literals can be created using double quotes or backticks. Double quotes are used to create strings that are human-readable and interpreted. Backticks are used to create raw string literals and are not interpreted.

```go
package main

import "fmt"

func main() {
    //String interpreted literal
    s1 := "Hello World"
    fmt.Println(s1)
    //String raw literal
    s2 := `Hello 
    World`
    fmt.Println(s2)

    //Other example
    fmt.Println("Hello World")
    fmt.Println(`Hello World`)
}
```

## String Concatenation

Strings can be concatenated using the `+` operator.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello" + " " + "World")
}
```

## String Length

The length of a string can be found using the `len()` function.

```go
package main

import "fmt"

func main() {
    fmt.Println(len("Hello World"))
}
```

## String Indexing

Strings can be indexed using the `[]` operator. The index starts at 0.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World"[1]) // 101
}
``` 

## String Slicing

Strings can be sliced using the `[:]` operator. The index starts at 0.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World"[1:5]) // ello
}
```

## String Formatting

Strings can be formatted using the `fmt.Sprintf()` function.

```go
package main

import "fmt"

func main() {
    fmt.Println(fmt.Sprintf("%v", 42)) // 42
    fmt.Println(fmt.Sprintf("%v", 42.42)) // 42.42
    fmt.Println(fmt.Sprintf("%v", true)) // true
    fmt.Println(fmt.Sprintf("%v", "Hello World")) // Hello World
}
```

## String Formatting with Type

Strings can be formatted with type using the `fmt.Sprintf()` function.

```go
package main

import "fmt"

func main() {
    fmt.Println(fmt.Sprintf("%T", 42)) // int
    fmt.Println(fmt.Sprintf("%T", 42.42)) // float64
    fmt.Println(fmt.Sprintf("%T", true)) // bool
    fmt.Println(fmt.Sprintf("%T", "Hello World")) // string
}
```

## String Formatting with Width

Strings can be formatted with width using the `fmt.Sprintf()` function.

```go
package main

import "fmt"

func main() {
    fmt.Println(fmt.Sprintf("%5v", 42)) //    42
    fmt.Println(fmt.Sprintf("%5v", 42.42)) // 42.42
    fmt.Println(fmt.Sprintf("%5v", true)) //  true
    fmt.Println(fmt.Sprintf("%5v", "Hello World")) // Hello World
}
```

## String Formatting with Precision

Strings can be formatted with precision using the `fmt.Sprintf()` function.

```go
package main

import "fmt"

func main() {
    fmt.Println(fmt.Sprintf("%.2v", 42)) // 42
    fmt.Println(fmt.Sprintf("%.2v", 42.42)) // 42.42
    fmt.Println(fmt.Sprintf("%.2v", true)) // true
    fmt.Println(fmt.Sprintf("%.2v", "Hello World")) // Hello World
}
```

## String Formatting with Padding

Strings can be formatted with padding using the `fmt.Sprintf()` function.

```go
package main

import "fmt"

func main() {
    fmt.Println(fmt.Sprintf("%05v", 42)) // 00042
    fmt.Println(fmt.Sprintf("%05v", 42.42)) // 42.42
    fmt.Println(fmt.Sprintf("%05v", true)) //  true
    fmt.Println(fmt.Sprintf("%05v", "Hello World")) // Hello World
}
```

## String Formatting with Padding and Precision

Strings can be formatted with padding and precision using the `fmt.Sprintf()` function.

```go
package main

import "fmt"

func main() {
    fmt.Println(fmt.Sprintf("%05.2v", 42)) // 42
    fmt.Println(fmt.Sprintf("%05.2v", 42.42)) // 42.42
    fmt.Println(fmt.Sprintf("%05.2v", true)) // true
    fmt.Println(fmt.Sprintf("%05.2v", "Hello World")) // Hello World
}
```

## String index unicode code point

Strings can be indexed using the `[]` %#U operator. The index starts at 0.

```go
package main

import "fmt"

func main() {
    a := "Hello World"
    for i, v := range a {
        fmt.Printf("index: %v, value: %v, unicode: %#U\n", i, v, v)
    }
}
    //print index: 0, value: 72, unicode: U+0048 'H'
    //print index: 1, value: 101, unicode: U+0065 'e'
    //print index: 2, value: 108, unicode: U+006C 'l'
    //print index: 3, value: 108, unicode: U+006C 'l'

>