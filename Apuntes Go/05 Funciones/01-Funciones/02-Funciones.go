package main

import "fmt"

// Función que retorna varios valores.

func calculosCuadrado(lado float64) (area float64, perimetro float64) {

	//En este caso devolvemos dos valores, area y perimetro. El orden será importante al recogerlos.

	perimetro = lado * 4
	area = lado * lado

	return area, perimetro // Devolvemos area y perimetro.
}

func main() {

	var lado float64

	fmt.Print("Inserte el valor del lado del cuadrado: ")
	fmt.Scanln(&lado)

	area, perimetro := calculosCuadrado(lado) // Recogemos el valor de las variables, teniendo en cuenta el orden en las que las hemos devuelto.

	fmt.Println("El perimetro del cuadrado es ", perimetro)
	fmt.Println("El area del cuadrado es ", area)

}
