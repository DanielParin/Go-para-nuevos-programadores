package main

import "fmt"

/*
	Las funciones variáticas nos sirven  para no delimitar el numero de parámetros que recibe la función,es decir,
	le podemos pasar los parámetros que queramos. Lo definimos con los ... antes del tipo de dato.
	Por ello como nos pueden pasar muchos numeros lo que hacemos es recorrerlo para que siempre opere con todos los datos.

	En nuestro caso, podremos pasar los numeros enteros que queramos, que los nombramos como numeros, y al recorrer numeros
	los vamos sumando. El bucle tendrá tantas iteraciones como numeros pasemos.
*/

func suma(numeros ...int) int {

	sumaTotal := 0

	for _, numero := range numeros {
		sumaTotal += numero
	}

	return sumaTotal
}

func main() {

	// Podemos pasar los números que queramos.
	fmt.Println(suma(5, 6))
	fmt.Println(suma(5, 6, 7, 8, 9))
}
