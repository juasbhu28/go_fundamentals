package main

import (
	"fmt"
)

/*
	El paquete fmt nos ayuda con muchas funciones para tipo texto
	crear archivos, leer archivos, imprimir en pantalla, etc.

	La documentación la encuentro en https://golang.org/pkg/fmt/ o godoc.org/fmt

	func Println(a ...interface{}) (n int, err error)
		Imprime en pantalla los valores que recibe, separados por un espacio
		y termina con un salto de linea.
		Sus parametros son de tipo interface{} que es un tipo generico que significa que puede
		recibir cualquier tipo de dato.
		retorna un entero que es la cantidad de bytes escritos y un error que es nil si no hay error.
*/

var a int
var b string = "Programa"
var c bool
var d bool = true

func main() {
	e := 42
	//String literal interpretado " " -> No permite agregar comillas simples ni saltos de linea
	f := "Dice hola mundo."
	// Raw string literal ` ` -> Permite imprimir saltos de linea
	g := `El doctor dicer que comer vegatles es
		 saludable.`
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("Mi", b, f)
	fmt.Println(g)

	//Vamos a usar printf para imprimir en pantalla
	//Hay tres patrones print, printf y println
	
	//Print -> Imprime en pantalla los valores que recibe, separados por un espacio
	//In go the verbs are:
	//%v	the value in a default format
	//	when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("El valor de la variable es: %v\n", a)
	//resultado es: El valor de la variable es: 0
	fmt.Printf("El tipo de a es: %T\n", a)
	//resultado es: El tipo de a es: int
	fmt.Printf("El valor de la variable es: %v\n", b)
	//resultado es: El valor de la variable es: Programa
	fmt.Printf("El tipo de b es: %T\n", b)
	//resultado es: El tipo de b es: string

	//The difference between %v and %s is that %v prints the value in the default format, and 
	//%s prints the value as a string.
	alias gplff="git pull --force-with-lease"

}

/* salida: 
0
Programa
false
true
42
Dice hola mundo.
Mi Programa Dice hola mundo.
El doctor dicer que comer vegatles es 
		 saludable.

*/

