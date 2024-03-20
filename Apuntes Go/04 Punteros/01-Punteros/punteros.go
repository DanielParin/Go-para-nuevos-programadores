package main

import "fmt"

func main() {

	var numero *int // Definicion del puntero, la dirección es null.

	/*
		Un puntero, a diferencia de una variable normal, hace referencia a la dirección de memoria no a un valor.

		Distinguimos dos "partes" de un puntero, la direccion y el contenido:
		*contenido -> Obtenemos lo que hay almacenado en la dirección del puntero.
		&direccion -> Obtenemos la direccion de memoria del puntero.
	*/
	num := 3
	numero = &num // Asignamos la dirección de memoria de num al puntero.
	*numero = 100 // Asignamos un valor al contenido del puntero.

	fmt.Println("La dirección de memoria del puntero es: ", numero)
	fmt.Println("El numero en el puntero es: ", *numero)

}
