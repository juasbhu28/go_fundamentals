# Ejercicio 1
# ==========

Usando el operador de declaración corta, asigna los siguientes valores a las variables con los identificadores "x", "y", y "z".

    a. 42
    b. "James Bond"
    c. true
Luego imprime los valores de cada variable.

    a. ¿Cuáles valores tienen las variables al imprimirse?
    b. ¿Qué tipo de variables son?

# Solución
```go
package main

import "fmt"

func main() {
    x := 42
    y := "James Bond"
    z := true
    fmt.Println(x, y, z)
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
    fmt.Printf("%T\n", x)
    fmt.Printf("%T\n", y)
    fmt.Printf("%T\n", z)
    //Result are:
    // 42 James Bond true
    // 42
    // James Bond
    // true
    // int
    // string
}

```

