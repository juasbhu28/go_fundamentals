/*
	En GO no se habla de Objetos
	Se crean TIPOS y VALORES
	No hay casting sino CONVERSIÓN y ASERCIÓN
*/

package main

import (
	"fmt"
)
//Declaro la variable con el identificador Z para que mantenga valores sólo de tipo int
var a int

type dinero int

var b dinero

func main() {	
	a = 42
	fmt.Println(a)
	fmt.Printlnf("%T\n", a)
		

	b = 42
	fmt.Println(a)
	fmt.Printlnf("%T\n", b)
	
	//Para CONVERTIR b a Int se usa la funcion int()
	b = int(b)
	fmt.Println(a)
	fmt.Printlnf("%T\n", b)

	/*
	Result 
		42
		int

		1000
		main.dinero

		1000
		int -> Aquí puede verse que b es un int y no de tipo dinero
	*/
}
