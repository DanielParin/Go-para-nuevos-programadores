package main

import (
	"errors"
	"fmt"
)

// Error global que usaremos luego
var divErr = errors.New("div errror, no 0")

func main() {

	var err error
	err = errors.New("mi error creado")
	fmt.Println(err)
	//Aquí imprimimos un error, no una cadena fmt.Println(err.Error())// Con err.Error() sacamos el error como string.

	//Vamos a probar distintas formas de resolver una divison entre 0 usando errores.

	result, error := div(2, 4) //Podríamos usar _, error para decir que la primera variable devuelta no la usaremos.

	//Devolviendo valor y viendo si es nulo
	if error == nil {
		fmt.Printf("El resultado es: %.2f\n", result)
	} else {
		fmt.Println(error)
	}

	//Devolviendo valor y usando errors.Is()
	if errors.Is(error, divErr) { // Si error es igual que divErr se cumple la condición.
		fmt.Println(error)
	} else {
		fmt.Println("El resultado es:", result)
	}

	//Panic()
	divPanic(2, 0)
	//El uso de panic no es recomendable, ya que dejas que el programa pete y no siga(probar a poner antes de los otros casos.
}

func divPanic(x, y float64) float64 {
	if y == 0 {
		panic("division by zero")
	} else {
		return x / y
	}
}

func div(x, y float64) (float64, error) {
	if y == 0 {
		return -1, divErr
	} else {
		return x / y, nil
	}
}
