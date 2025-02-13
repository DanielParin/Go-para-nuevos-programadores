package main

import "fmt"

/*
	Las estructuras son una forma de agrupar varios datos en una sola variable que contenga todos ellos.
	Podemos definir una estructura con todos los campos que queramos y del tipo que queramos.
	Personalmente, algo que me ayudó a entenderlo fue ver una estructura como un tipo de dato personalizado.
	En vez de tener un string o int, creamos un tipo de dato que engloba más datos.

*/

// Definición de la estructura
type Estructura struct {
	Campo1 string
	Campo2 int
	Campo3 bool
	Campo4 []string
}

func main() {

	// Instanciamos una variable del tipo Estructura.
	variableEstructura := Estructura{
		Campo1: "Nombre del campo",
		Campo2: 1,
		Campo3: true,
		Campo4: []string{"Array", "de", "la", "estructura."},
		// Si no hubiéramos dado un valor a algún campo se habría quedado sin valor (int -> 0, string -> "") pero no salta error.
	}
	fmt.Printf("%+v", variableEstructura) // El %+v nos sirve para que nos imprima el nombre del campo y su valor, si ponemos solo la variable nos imprime los valores.

	// Para acceder a un valor de la variable: Ponemos el <nombreVariableEstructura>.<campoAMostrar>
	fmt.Println("\nEl campo 2 es: ", variableEstructura.Campo2)

	/*
		Otra forma de instanciar una estructura es usando new(), lo que internamente crea la estructura y reserva la memoria para la variable,
		y pone todos los campos sin valor.
		Lo bueno de esto es que siempre tendremos espacio para dar valores a la estructura y poder usarla.
		Luego podemos darle valores a los campos o lo que queramos.
	*/

	variableEstructuraNew := new(Estructura)
	fmt.Printf("%+v", variableEstructuraNew)
	fmt.Println()

	/*
		Si venimos de un lenguaje orientado a objetos podemos asignar valores de una forma parecida a como se hace en ese tipo de lenguajes.
		Usando <nomEstructura {<valorCammpo>,<valorCampo>......}
		Lo único a tener en cuenta es que si no quisiéramos indicar un valor, no podríamos usar esta forma.
	*/

	variableEstructura3 := Estructura{"ABCD", 2, true, []string{"Array", "estructura"}}

	fmt.Printf("%+v", variableEstructura3)

}
