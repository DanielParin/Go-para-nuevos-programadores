package main

import "fmt"

func saludar() {
	fmt.Println("Holaaaa")
}

func despedirse() {
	fmt.Println("Adios")
}

func hablar() {
	fmt.Println("BlaBlaBla...")
}

/*
	Lo que hace el defer es hacer que una función se ejecute al final de todo el código donde esté definida, aunque secuencialmente esté antes
	de otras funciones (En los casos del main se puede ver perfectamente).
	En caso de tener varias funciones con defer, se ejecutarán en orden contrario al leido, es decir, la primera funcion con defer leida será
	la última en ejecutarse.
*/

func main() {

	// Ejecucion sin defer
	saludar() // Probar a poner defer aquí.
	despedirse()
	hablar()
	hablar()
	hablar()

	fmt.Println("\n\n") // Saltos de linea para diferenciar

	saludar()
	defer despedirse()
	hablar()
	hablar()
	hablar()
}
