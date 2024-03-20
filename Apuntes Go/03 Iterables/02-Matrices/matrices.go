package main

import "fmt"

func main() {

	var matriz [3][3]int
	var numero int

	/*
		Una matriz es igual que un array(a nivel de operar con el) pero implica dos dimensiones.
		Mientras que un array es: [1 2 3 4.....].
		Una matriz es:
						[1 2 3 4]
						[5 6 7 8]
						[9 0 1 2]
						........
	*/

	// Vamos a rellenar la matriz con un numero y sus siguientes.

	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&numero)

	//Para rellenar una matriz necesitamos dos for, uno para la fila y otro para la columna
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = numero
			numero += 1
		}
	}

	fmt.Println(matriz) //Esto nos pinta la matriz
	fmt.Print("\n\n\n")

	//Asi vamos a pintar la matriz mas "bonita"
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println() //Para hacer un salto de linea despues de cada fila
	}
}
