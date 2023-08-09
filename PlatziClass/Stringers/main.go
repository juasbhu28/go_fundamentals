package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Stringer is an interface that defines the String method
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)
}

//The result of stringer is an output customized
