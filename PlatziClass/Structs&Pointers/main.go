package main

import (
	"fmt"
)

//Structs & Pointers
//Structs is a type of data that allows us to group values of different types of data
//Structs are like classes in other languages
// & is used to get the memory address of a variable
// * is used to get the value of a memory address

type User struct {
	edad             int
	nombre, apellido string
}

func main() {

	a := 50
	b := &a //b is a pointer to a

	fmt.Println("Print the value of a")
	fmt.Println(a)
	fmt.Println("Print the memory address of a")
	fmt.Println(&a)
	fmt.Println("Print the value of b")
	fmt.Println(b) //Prints the memory address of a
	fmt.Println("Print the value of *b")
	fmt.Println(*b) //Prints the value of a

	var objeto User
	objeto.edad = 23
	objeto.nombre = "Juan"
	objeto.apellido = "Perez"

	fmt.Println("Print the value of the struct")
	fmt.Println(objeto)
	fmt.Println("Print the value of the struct using the dot notation")
	pointer := &objeto
	fmt.Println(pointer) // Prints the memory address of the pointer
	fmt.Println("Print the value of the pointer using the *")
	fmt.Println(*pointer) //Prints the value of the pointer

	//If we change the value of the pointer, the value of the object will change too
	pointer.nombre = "Carlos"
	fmt.Println(*pointer) //Prints the value of the pointer
	fmt.Println(objeto)   // Prints the value of the object

	//Structs anonimos
	anonimo := struct {
		edad               int
		nombre, apellido   string
		estudiante, casado bool
	}{
		edad:       23,
		nombre:     "Juan",
		apellido:   "Perez",
		estudiante: true,
		casado:     false,
	}

	fmt.Println(anonimo)

	//Structs anonimos can be used as a parameter of a function
	//This struct is only available inside the function
	//This is a good practice to avoid creating a lot of structs

	//The next fuction is an anonymous function, it is a function without a name
	//It is used when we need to create a function that will be used only once
	//In this case is used to create an anonymous struct inside the function and print it
	func() {
		anonimo := struct {
			edad               int
			nombre, apellido   string
			estudiante, casado bool
		}{
			edad:       23,
			nombre:     "Juan",
			apellido:   "Perez",
			estudiante: true,
			casado:     false,
		}

		fmt.Println(anonimo)
	}()
	//Using Person Struct and methods on and pointers
	var myPerson Person
	myPerson.Edad = 23
	myPerson.Nombre = "Juan"
	myPerson.Apellido = "Peres"
	fmt.Print(myPerson.Init(myPerson))
	fmt.Println(myPerson.PrintUnaVez())
	myPerson.PrintPuntero() //This will change the value of the struct
	fmt.Println(myPerson.PrintUnaVez())
}
