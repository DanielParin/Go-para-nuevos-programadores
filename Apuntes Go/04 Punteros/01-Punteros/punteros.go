package main

import "fmt"

func main() {

	var numero *int // Definición del puntero, la dirección es null.

	/*
		Un puntero, a diferencia de una variable normal, hace referencia a la dirección de memoria no a un valor.

		Distinguimos dos "partes" de un puntero, la dirección y el contenido:
		*contenido -> Obtenemos lo que hay almacenado en la dirección del puntero.
		&dirección -> Obtenemos la dirección de memoria del puntero.
	*/
	num := 3
	numero = &num // Asignamos la dirección de memoria de num al puntero.
	*numero = 100 // Asignamos un valor al contenido del puntero.

	fmt.Println("La dirección de memoria del puntero es: ", numero)
	fmt.Println("El numero en el puntero es: ", *numero)

	/*
		Como el puntero toca directamente la memoria, es peligroso y útil, ya que podemos modificar su valor externamente
		sin necesidad de devolvérselo y asignárselo.
	*/

	var valorNumero int
	valorNumero = 10

	var valorPuntero = new(int) //Definimos la variable que usaremos como puntero y lo instanciamos para reservar una dirección de memoria
	*valorPuntero = 20          //Asignamos el valor contenido a 20.

	fmt.Println("El numero antes de la función vale: ", valorNumero)
	fmt.Println("El puntero antes de la función vale: ", *valorPuntero)

	incrementarNumero(valorNumero)
	incrementarPuntero(valorPuntero)

	fmt.Println("El numero después de la función vale: ", valorNumero)
	fmt.Println("El puntero después de la función vale: ", *valorPuntero)
	//Vemos que al usar punteros si se modifica el valor, porque no tocamos una variable local, sino que tocamos la dirección de memoria.

}

func incrementarNumero(valor int) {
	valor++
}

func incrementarPuntero(valor *int) {
	*valor++
}
