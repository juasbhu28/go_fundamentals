package main

import "fmt"

type figura interface {
	area() float64
}

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figura) {
	fmt.Println("Area: ", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}
	calcular(myCuadrado)
	calcular(myRectangulo)
}
