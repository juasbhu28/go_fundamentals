package main

import "fmt"

//El operador de claración corto solo permite usar la variable dentro
//del escope donde se creo a excepción de usar la palabra reservada VAR para delcararlo

var z = 41

func main() {	
	x := 42 + 7
	y := "String"
	fmt.Println(x)	
	fmt.Println(y)
	x = 50
	fmt.Println(x, y)
	fmt.Println(z)
	numero()
}

func numero(){
	//Aqui puedo usar la variable z por qué esta a nivel global con la palabra var
	fmt.Println(z)
}