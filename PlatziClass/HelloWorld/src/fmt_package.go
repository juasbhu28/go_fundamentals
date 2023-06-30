package main

import "fmt"

func main() {
	//Variables
	helloMessage := "Hello"
	worldMessage := "World"

	//Println
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	//Terminal output: Platzi tiene más de 500 cursos
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos) // %v es un comodín que se puede usar para cualquier tipo de dato en caso de
		//que no se sepa el tipo de dato que se va a imprimir

	//Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message) //Sprintf regresa un string con el formato que se le haya dado

	

}