package main

import "fmt"

/*
	Las funciones se usan para ordenar el código.
	Para entenderlo facilmente, lo que hasta ahora haciamos en el main lo podemos ir separando en distintas funciones.
	Las funciones pueden recibir parámetros o no, que serán variables/valores que enviamos desde el main u otras funciones.
	Las funciones básicas pueden ser de dos tipos:
	- Que retornen algo: Son funciones que al final tendran un return para devolver un valor
	- Que no retorne nada:  Son funciones que no devuelven nada.
*/

// Función que retorna un valor.
func calcularPerimetro(lado float64) float64 {

	/*
		En esta funcion recibiremos un float64 al que llamamos lado. No tiene porque llamarse igual que la variable que enviamos, ya que
		simplemente como la llamemos será asi para dentro de la función, pero si que tiene que tener el mismo tipo de dato.
	*/
	return lado * 4 // Aquí retornamos el valor de la operación que tiene que ser del mismo tipo que el asignado.
}

// Funcion que no retorna nada.
func calcularArea(lado float64) {

	//En este caso simplemente recibimos otro float64 al que llamamos lado.

	fmt.Println("El area es ", lado*lado)
}
func main() {

	var lado float64
	var perimetro float64
	fmt.Print("Inserte el valor del lado del cuadrado: ")
	fmt.Scanln(&lado)

	perimetro = calcularPerimetro(lado) // Enviamos lado como parametro y lo que nos retorne lo almacenamos en perimetro.
	fmt.Println("El perimetro es ", perimetro)

	calcularArea(lado) // Enviamos lado como parametro y como no retorna nada, no lo almacenamos en otra variable.
}
