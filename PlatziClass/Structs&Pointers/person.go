package main

import "fmt"

//Returning a structs in fucntions

type Person struct {
	Edad             int
	Nombre, Apellido string
}

func (myPerson Person) Init(p Person) Person {
	myPerson.Edad = p.Edad
	myPerson.Nombre = p.Nombre
	myPerson.Apellido = p.Apellido
	return myPerson
}

// When we use (this Person) we are saying that the function is a method of the Person struct
// And we are saying that the function is not going to change the value of the struct
// This is a good practice to avoid changing the value of the struct
// And to avoid creating a copy of the struct
// Finally, returning a String
func (myPerson Person) PrintUnaVez() string {
	myPerson.Nombre = "Carlos" //This will not change the value of the struct
	return myPerson.Nombre + " " + myPerson.Apellido
}

// When we use (this *Person) we are saying that the function is a method of the Person struct
// And we are saying that the function is going to change the value of the struct
// This is a good practice to avoid creating a copy of the struct
// Finally, we print the value of the struct
func (myPersonPointer *Person) PrintPuntero() {
	myPersonPointer.Nombre = "Carlos" //This will change the value of the struct
	fmt.Println(myPersonPointer.Nombre + " " + myPersonPointer.Apellido)
}
