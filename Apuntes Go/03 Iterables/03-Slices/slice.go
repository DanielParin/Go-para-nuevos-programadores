package main

import "fmt"

func main() {

	/*
		Un slice es parecido a un array, lo único que tiene tamaño variable y suele hacer referencia a un array o coger datos de este.
		Al igual que el array, es un conjunto de datos de un mismo tipo.
	*/
	arrayNumeros := []int{1, 2, 3, 4, 5, 6} //Array de numeros enteros.
	lenguajes := make([]string, 2)          // Creamos un slice de strings de tamaño 2.

	lenguajes[0] = "Go"   // La posición 0 tendrá el valor Go
	lenguajes[1] = "Java" // La posición 1 tendrá el valor Java

	fmt.Println("Slice de lenguajes: ", lenguajes)

	lenguajes = append(lenguajes, "Kotlin", "C") // Añadimos dos elementos al slice
	fmt.Println("Slice de lenguajes: ", lenguajes)

	sliceNumeros := arrayNumeros[1:4] // Creamos un slice a partir de un array (coge las posiciones 1,2 y 3).
	fmt.Println("Slice de numeros: ", sliceNumeros)

	sliceNumeros[0] = 9                     // Cambiamos el primer valor del slice.
	sliceNumeros = append(sliceNumeros, 10) // Añadimos otro valor
	fmt.Println("Slice de numeros: ", sliceNumeros)

	//Eliminar elemento del slice

	indiceBorrar := 3 // El indice que vamos a borrar

	// Aqui copiamos hasta el indice que queremos borrar y nos saltamos ese elemento, y seguimos copiando desde el siguiente índice.
	lenguajes = append(lenguajes[:indiceBorrar], lenguajes[indiceBorrar+1:]...)
	fmt.Println("Slice sin 3 indice: ", lenguajes)

	// Copiar un slice en otro slice

	// Creamos un nuevo slice
	var lenguajesCopy = make([]string, len(lenguajes)) // Creamos un slice con la longitud
	copy(lenguajesCopy, lenguajes)                     // Copiamos  el contenido de lenguajes en lenguajesCopy.

	fmt.Println("Slice copiado: ", lenguajesCopy)
}
