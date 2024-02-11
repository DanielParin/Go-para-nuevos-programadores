package main

import "fmt"

func main() {

	var numeroUno int
	var numeroDos int
	var resultado float64 //resultado decimal
	var eleccion int
	resultadoObtenido := true //Lo iniciamos a true para que se presuponga que habrá valor

	//Primero seleccionamos los valores de los numeros con los que vamos a operar.
	fmt.Print("Seleccione un numero: ")
	fmt.Scanln(&numeroUno)
	fmt.Print("Seleccione otro numero: ")
	fmt.Scanln(&numeroDos)

	fmt.Println("1. Suma")
	fmt.Println("2. Resta")
	fmt.Println("3. Multiplicacion")
	fmt.Println("4. Division")
	fmt.Print("Elija operacion: ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case 1:
		resultado = float64(numeroUno) + float64(numeroDos) //Parseamo los numeros para que de enteros pasen a ser decimales.
	case 2:
		resultado = float64(numeroUno) - float64(numeroDos)
	case 3:
		resultado = float64(numeroUno) * float64(numeroDos)
	case 4:
		resultado = float64(numeroUno) / float64(numeroDos)
	default:
		resultadoObtenido = false //Para que luego no nos imprima si no hay resultado.
		fmt.Println("Operacion no encontrada.")
	}

	if resultadoObtenido {
		fmt.Printf("El resultado es %f", resultado)
	}

}
