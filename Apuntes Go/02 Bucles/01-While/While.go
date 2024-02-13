package main

import "fmt"

func main() {

	var numeroElegido int

	fmt.Print("Inserte numero: ")
	fmt.Scanln(&numeroElegido)

	for numeroElegido != 0 {

		fmt.Print("Inserte numero: ")
		fmt.Scanln(&numeroElegido)
	}
	fmt.Println("Has insertado 0.")
}

/*
	El bucle while es distinto al de otros lenguajes, en Go no tenemos ningún bucle while como tal, si no que tenemos una variación del bucle for pero al
	fin y al cabo implican lo mismo.
	Es un bucle indefinido, que quiere decir esto, que no sabemos cuantas veces se va a ejecutar, puede ejecutarse una vez o infinitas, su funcionalidad
	es simple, mientras que se cumpla la condición solo se ejecutará el codigo de dentro hasta que no se cumpla y entonces salga del bucle.
	En nuestro caso no sabemos cuando la persona que ejecute el programa insertará un cero, asique mientras que el numero insertado no sea cero, nos seguirá pidiendo
	hasta que sea cero y ya se salga del bucle.
*/
