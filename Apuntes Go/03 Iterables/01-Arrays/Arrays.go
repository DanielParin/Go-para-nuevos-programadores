package main

import "fmt"

/*
	Entendemos como array a un conjunto de datos del mismo tipo. Para acceder a esos datos usamos los índices del array, por así decirlo,
	la posición del dato. LOS INDICES EMPIEZAN POR 0 NO POR UNO.
	Un array de tamaño 3 si nos lo imaginaramos sería  array1 = [índice0,índice1,índice2] esos datos solo podrán ser de un mismo tipo.
*/

func main() {

	var array1 = [3]int{1, 2, 3} //Primera forma de definirlo, con  var nombreArray =[tamaño]tipo{valores}

	array2 := [3]int{} // Segunda forma de definirlo, con  nombreArray:=[tamaño]tipo{valores}

	fmt.Println(array1) //Imprimimos el array

	//Para rellenar los arrays podemos usar un for, asi podemos recorrer todas las posiciones y agregar valor a cada una

	for i := 0; i < 3; i++ {
		fmt.Print("Seleccione un numero: ")
		fmt.Scanln(&array2[i]) //La primera iteración guarda el valor en la posicion 0 y asi continuamente.
	}
	fmt.Println(array2)

	//Para usar los valores del array, igualmente se usan las posiciones.

	array3 := [3]int{}

	fmt.Println("Suma de los dos arrays.")
	for i := 0; i < 3; i++ {
		//array3 = array1 + array2 // Se tiene que sumar posicion a posicion
		array3[i] = array1[i] + array2[i]
	}
	fmt.Print(array3)

}
