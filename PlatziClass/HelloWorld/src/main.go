package main

import (
	"fmt"

	"github.com/go_fundamentals/PlatziClass/HelloWorld/src/acces/private"
	"github.com/go_fundamentals/PlatziClass/HelloWorld/src/acces/public"
)

func mainTwo() {
	fmt.Println("Hello World")

	//Variables

	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println(pi)
	fmt.Println(pi2)

	//Declaración de variables enteras
	base := 12          // := solo se usa cuando se declara la variable por primera vez
	var altura int = 14 // var se usa cuando se declara la variable por primera vez y se le asigna un tipo de dato y valor
	var area int        // la creamos pero no le asignamos valor

	//GO es delicado con el uso de varaibles que no se usan

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es: ", areaCuadrado)

	//Operadores aritmeticos
	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	//División
	result = y / x
	fmt.Println("División: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	//Calcular el area de un rectangulo, trapecio y circulo
	//Rectangulo
	baseRectangulo := 20
	alturaRectangulo := 10
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("El area del rectangulo es: ", areaRectangulo)

	//Trapecio
	baseMayorTrapecio := 20
	baseMenorTrapecio := 10
	alturaTrapecio := 10
	areaTrapecio := ((baseMayorTrapecio + baseMenorTrapecio) * alturaTrapecio) / 2
	fmt.Println("El area del trapecio es: ", areaTrapecio)

	//Circulo
	const piCirculo float64 = 3.14
	radioCirculo := 10
	areaCirculo := piCirculo * float64(radioCirculo*radioCirculo)
	fmt.Println("El area del circulo es: ", areaCirculo)

}

func usingPackagePublic() {
	var myCar public.CarPublic
	myCar.PublicBrand = "Ferrari"
	myCar.PublicYear = 2020

	//En un mismo archivo puedo crear diferentes structs, publicos y privados

}

func usingPackagePrivate() {
	// var myCar privat
	var myCar public.CarPublic
	myCar.PublicBrand = "Ferrari"
	myCar.PublicYear = 2020

	//Para acceder a un paquete privado desde otro paquete, debe existir una función
	//Que permita acceder a ese paquete privado
	private.PrintCarPrivate(myCar)

}
