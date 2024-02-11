package main

import "fmt"

func main() {

	var respuesta int

	fmt.Print("Seleccione un numero: ")
	fmt.Scanln(&respuesta)

	switch respuesta {

	case 1:
		fmt.Println("Has seleccionado 1.")
	case 2:
		fmt.Println("Has seleccionado 2.")
	case 3:
		fmt.Println("Has seleccionado 3.")
	default:
		fmt.Println("El numero seleccionado no esta entre los 3 primeros.")
	}

	/*
		El switch es como una cadena de else-if pero de una manera más organizada.

		Se basa en el valor de una variable y dependeiendo del valor se ejecuta el case asociado a ese valor(el valor es lo que va despues del case), si no hay
		ningún case asociado a ese valor se ejecutará el default.

		En nuestro caso, si insertamos un 1 por consola se ejecutará el case 1 y si insertamos un valor distinto a 1,2 y 3 se ejecuta el default.
	*/
}
