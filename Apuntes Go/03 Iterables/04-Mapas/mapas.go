package main

import "fmt"

/*
	Los mapas tambien son iterables, pero, en vez de contener un conjunto de elementos del mismo tipo contienen pares
	clave-valor. Es decir el mapa contiene un elemento llamado clave que se asocia con su valor.
	Clave: 1
	Valor: Uno
	En este caso la clave 1 esta asociada con el valor uno.
*/

func main() {

	// Creamos el mapa que sera String-String
	var mapaPersonas = make(map[string]string) // map[clave]valor

	mapaPersonas["46253895T"] = "Carlos Raul" // Creamos un elemento del mapa indicando una nueva clave y un valor
	fmt.Println(mapaPersonas)

	mapaPersonas["46253895T"] = "Carlos Raul Jimenez" // Si intentamos crear un nuevo elemento y la clave existe, modifica solo el valor
	fmt.Println(mapaPersonas)

	// Para acceder al valor de la clave:
	fmt.Println(mapaPersonas["46253895T"])
	fmt.Println("\n\n")

	// Para eliminar un elemento
	mapaPersonas["33333333T"] = "Persona a eliminar"
	fmt.Println(mapaPersonas)

	delete(mapaPersonas, "33333333T") // Para borrar un elemento indicamos la clave del par a borrar.
	fmt.Println(mapaPersonas)         // Veremos que ha eliminado el par clave-valor indicado.
}
