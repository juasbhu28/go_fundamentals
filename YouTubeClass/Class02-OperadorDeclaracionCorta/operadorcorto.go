package main

import "fmt"

//Todos las estructuras en go implementan la interface vacia.
//Ejemplo fmt.Println("Hello, 世界", 45, true)


//No puedo usar el operador corto fuera del scope de la función.


func main() {
	x := 42
	fmt.Println(x)

	// Go usa operadores y comandos cada lina es (Declaraciones y expresiones)
	// Statements and declarations
	y := x + 8

	//En go no puedo cambiar el valor de una variable ya existente con el operador
	//de declaración corta, necesito usar "="
	x = 9
	
	fmt.Println(x, y)
	

}
