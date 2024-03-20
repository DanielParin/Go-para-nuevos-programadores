package main

import "fmt"

/*
	Las funciones recursivas son un tipo de funcion que se caracteriza porque en el código definido en la funcion se vuelve a llamar así misma.
	Esto nos daría un ciclo infinito ya que  la propia funcion se llama así misma todo el rato, por eso tienes que hacer una condición de salida,
	es decir, el punto donde acaba/ retorna un valor y ya no se llama más asi misma.
*/

func factorial(numero int) int {

	// Condicion de salida
	if numero <= 1 {
		return numero
	}

	// Recursividad y es lo que retornará al final.
	return numero * factorial(numero-1)

}

func main() {

	var numero int
	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&numero)

	//Imprimimos el resultado final.
	fmt.Println(factorial(numero))
}
