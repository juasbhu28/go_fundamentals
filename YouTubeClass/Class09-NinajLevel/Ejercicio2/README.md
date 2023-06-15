#  \ escape characters

[]: # Path: Class09-NinajLevel/Ejercicio2/README.md

Adding new lines in the print function to make it more readable using the \n

We could use another slahs for other characters like 

    - \t for tabulation
    - \b for backspace
    - \r for carriage return
    - \a for alert
    - \v for vertical tabulation
    - \f for form feed
    - \n for new line
    
## Ejercicio 2

Usar var para declarar tres variables. Las variables deben tener scope a nivel de paquete. No asignar valores a las variables. Usa los siguientes identificadores para las variables y asegúrate de que las variables son de los siguientes tipos en go usa el operador typeof para determinar el tipo de una variable.

    - identificador “x” tipo int
    - identificador “y” tipo string
    - identificador “z” tipo bool

En main

1. imprime los valores de cada identificador
2. el compilador asigna un valor para cada identificador. ¿Cuál es el valor asignado para cada identificador?
3. ahora usa conversión para convertir los valores de tipo int y bool a tipo string e imprime los valores convertidos.

## Solución
    
```go

        package main
        
        import "fmt"
        
        var x int
        var y string
        var z bool
        
        func main() {
        	fmt.Println(x)           // 0-> valor por defecto
        	fmt.Println(y)           // <nil> -> valor por defecto
        	fmt.Println(z)           // false -> valor por defecto
        	fmt.Printf("%T\n", x)    // int -> tipo de dato
        	fmt.Printf("%T\n", y)    // string -> tipo de dato
        	fmt.Printf("%T\n", z)    // bool -> tipo de dato
        	x = 42
        	y = "James Bond"
        	z = true
        	fmt.Println(x)
        	fmt.Println(y)
        	fmt.Println(z)
        	fmt.Printf("%T\n", x)
        	fmt.Printf("%T\n", y)
        	fmt.Printf("%T\n", z)
        	xs := fmt.Sprintf("%v", x) // convertir a string
        	ys := fmt.Sprintf("%v", y)
        	zs := fmt.Sprintf("%v", z)
        	fmt.Println(xs)            // 42
        	fmt.Println(ys)
        	fmt.Println(zs)
        	fmt.Printf("%T\n", xs)     // string
        	fmt.Printf("%T\n", ys)     // string
        	fmt.Printf("%T\n", zs)     // string
        }

```

### Resultado
    
        0
        <nil>
        false
        int
        string
        bool
        42
        James Bond
        true
        int
        string
        bool
        42
        James Bond
        true
        string
        string
        string