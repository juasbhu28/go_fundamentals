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

	//Múltiples condiciones anidadas con switch
	//Número par o impar
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}

	//Keyword defer, se ejecuta al final de la función. Antes de finalizar la función se ejecuta el defer
	//Se puede usar para cerrar archivos, conexiones a bases de datos, etc.
	fmt.Println("Conexión a la base de datos")
	defer fmt.Println("Cerrar la base de datos") //Ultima función antes de finalizar la función main
	fmt.Println("Consultamos información, set de datos")

	//Puedes usar más de una vez defer pero la buena práctica es usarlo una sola vez al final de la función

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Es 2")
			continue //Salta a la siguiente iteración
		}
		if i == 8 {
			fmt.Println("Break")
			break //Termina el ciclo
		}
	}

	//Arrays y Slices
	var array [4]int
	fmt.Println(array) //Imprime [0 0 0 0]

	array[0] = 1
	array[1] = 2

	fmt.Println(array) //Imprime [1 2 0 0]

	totalElementos := len(array)
	fmt.Println("Tamaño del arreglo:", totalElementos) //Imprime 4

	capacityArray := cap(array)
	fmt.Println("Capacidad del arreglo:", capacityArray) //Imprime 4

	//Slices
	slice := []int{1, 2, 3, 4, 5, 6} // ->No indicamos la cantidad de elementos que tendrá el slice
	fmt.Println(slice)               //Imprime [1 2 3 4 5 6]
	fmt.Println("Tamaño del slice:", len(slice))
	fmt.Println("Capacidad del slice:", cap(slice))

	//Métodos en el slice
	fmt.Println(slice[0])   //Imprime 1
	fmt.Println(slice[:3])  //Imprime [1 2 3]
	fmt.Println(slice[2:4]) //Imprime [3 4]
	fmt.Println(slice[4:])  //Imprime [5 6]

	//Slice con make -> Función que nos permite crear slices
	sliceMake := make([]int, 3, 5) // ->Indicamos la cantidad de elementos que tendrá el slice
	fmt.Println(sliceMake)         //Imprime [0 0 0]
	fmt.Println("Tamaño del sliceMake:", len(sliceMake))
	fmt.Println("Capacidad del sliceMake:", cap(sliceMake))

	//Los arrays son inmutables, los slices son mutables.
	//Los slices son referencias a un array, si modificamos un slice, modificamos el array original

	//Append
	sliceAppend := []string{"Hola", "Mundo", "Go", "Platzi"}
	fmt.Println(sliceAppend) //Imprime [Hola Mundo Go Platzi]
	sliceAppend = append(sliceAppend, "Curso")
	fmt.Println(sliceAppend) //Imprime [Hola Mundo Go Platzi Curso]

	//Append add new list
	newSlice := []string{"Welcome", "to", "the", "course"}
	sliceAppend = append(sliceAppend, newSlice...) // ... -> Indica que se va a agregar un nuevo slice

	//Copy
	sliceCopy := make([]string, len(sliceAppend))
	copy(sliceCopy, sliceAppend)
	fmt.Println(sliceCopy)

	//Slices con Range
	//Primer valor es el indice y el segundo es el valor
	for i, v := range sliceAppend {
		fmt.Println(i, v)
	}

	//Si no nos interesa el valor del indice podemos usar _
	for _, v := range sliceAppend {
		fmt.Println(v)
	}

	//Si solo necesitamos el indice
	for i := range sliceAppend {
		fmt.Println(i)
	}

	//El range es un iterador, si no se usa el valor del indice o el valor, se puede usar un guión bajo

	//Palindromo con range y slices GO
	//anita lava la tina

	frase := "anita lava la tina"
	isPalindromo(frase)

	learningMaps()

}

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i]) //La función string convierte el valor a string porque el valor de text[i] es un entero ascii
	}
	if textReverse == text {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func learningMaps() {
	//Maps
	//Los maps son una estructura de datos que nos permite almacenar valores asociados a una llave
	//La llave puede ser de cualquier tipo que se pueda comparar con el operador ==
	//Los valores pueden ser de cualquier tipo y pueden ser diferentes entre sí

	//Declaración de un map
	//map[Tipo de la llave]Tipo del valor
	m := make(map[string]int)

	//Agregar elementos al map
	m["Jose"] = 14
	m["Pedro"] = 20
	m["Juan"] = 30

	fmt.Println(m) //Imprime map[Jose:14 Pedro:20 Juan:30]

	//Obtener un valor del map
	fmt.Println(m["Jose"]) //Imprime 14

	//Obtener un valor que no existe en el map
	fmt.Println(m["Luis"]) //Imprime 0

	//Verificar si una llave existe en el map
	//El Ok nos indica si la llave está en el diccionario
	value, ok := m["Luis"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("La llave no existe en el map")
	}

	//Eliminar un elemento del map
	delete(m, "Pedro")
	fmt.Println(m) //Imprime map[Jose:14 Juan:30]

	//Recorrer un map
	for key, value := range m {
		fmt.Println(key, value)
	}
}
