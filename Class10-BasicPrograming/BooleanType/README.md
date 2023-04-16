# Boolean type

    A boolean type is a data type that has only two possible values: true and false in go
    A differents others languages when null is a value valid, in go is not valid

    In go, the zero value of a boolean type is false

    Example:

    ```go

    package main

    import "fmt"

    var x bool

    func main() {
        fmt.Println(x) // zero value -> false
        x = true
        fmt.Println(x) // true
    }

    ```

output:

    false
    true

## Boolean type with if statement

    In go, the if statement is a control flow statement that allows to execute a block of code if a condition is true

    Operators:
    
    - == equal
    - <= less than or equal
    - >= greater than or equal
    - != not equal
    - > greater than
    - < less than

## Algebraic boolean operators

    In go, the algebraic boolean operators are used to combine boolean expressions

    Operators:
    
    - && and
    - || or
    - ! not

    Example:

    ```go

    package main

    import "fmt"

    func main() {
        x := 42
        y := 42

        fmt.Println(x == y) // true
        fmt.Println(x <= y) // true
        fmt.Println(x >= y) // true
        fmt.Println(x != y) // false
        fmt.Println(x > y) // false
        fmt.Println(x < y) // false

        fmt.Println(x == 42 && y == 42) // true
        fmt.Println(x == 42 || y == 42) // true
        fmt.Println(!(x == 42)) // false
    }

    ```