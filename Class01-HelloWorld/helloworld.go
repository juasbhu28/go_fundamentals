//Todos los paquetes de go deben tener un paquete
package main

//La forma en que importo paquetes
/*
* import (
*	"paquete1"
*	"paquete2"
*  )
*
*/
import "fmt"

//Cuando una función empieza por mayuscula son de tipo "publico" es decir que permiten 
//exportarse a otras librerias

/*Varaible parameter ("...") Es decir que recibe un número de variables no finitas
de un tipo especifico, en el ejemplo de tipo interface provienen todas las clases entonces recibe cualquier tipo
Ejemplo 
	func Println(a ...interface{})


Go no permite usar variables que no se usen a excepción de cuando inicia con raya al piso " _ "

:= Esto es el operador de declaración de variables corto para declarar nuevas variables
=
*/


func main() {
	nb, err := fmt.Println("Hello, 世界")
	_, _ = fmt.Println(nb, err)
}

