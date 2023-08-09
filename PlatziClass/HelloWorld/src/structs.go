//Structs es un tipo de dato que nos permite agrupar valores de diferentes tipos de datos

package main

import "fmt"

type car struct {
	brand string
	year  int
}

func mainStruct() {
	myCar := car{brand: "Ford", year: 2020}
	println(myCar)
	println(myCar.brand)
	println(myCar.year)
	//Otra forma de declarar un struct o instanaciar
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}

// Funcion que retorna un struct
func myCar(brand string, year int) car {
	return car{brand, year}
}

// Funcion que retorna un puntero a un struct
func myCarPointer(brand string, year int) *car {
	return &car{brand, year}
}

// Funcion que retorna una lista de structs
func myCars() []car {
	cars := []car{{"Ford", 2020}, {"Ferrari", 2021}}
	return cars
}
