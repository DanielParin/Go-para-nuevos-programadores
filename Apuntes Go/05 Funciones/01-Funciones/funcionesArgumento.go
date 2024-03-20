package main

import "fmt"

/*
En este caso lo que devolvemos son funciones, al estar definidas y usadas en el mismo ámbito, las llamamos funciones anónimas.
Hay que tener en cuenta que luego cuando recojamos lo que devuelve no tenemos el valor devuelto por area, si no la propia
función area, entonces al usarla ponemos area().
En el argumento ponemos: nombreFuncion func() valorDevuelto , (más funciones si fueran necesarias)
*/
func calculoCuadrado(lado float64) (area func() float64, perimetro func() float64) {

	area = func() float64 {
		return lado * lado
	}
	perimetro = func() float64 {
		return lado * 4
	}
	return
}
func main() {

	var lado float64
	fmt.Print("Inserte la longitud del lado: ")
	fmt.Scanln(&lado)

	area, perimetro := calculoCuadrado(lado) // Recogemos las funciones que devuelve.

	fmt.Println("El area del cuadrado es ", area())           // area() es la funcion recogida en la variable area.
	fmt.Println("El perimetro del cuadrado es ", perimetro()) // perimetro() es la funcion recogida en la variable perimetro.

}
