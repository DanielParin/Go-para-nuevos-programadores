package main

import "fmt"

func main() {

	var numeroElegido int

	for {

		fmt.Print("Inserte numero: ")
		fmt.Scanln(&numeroElegido)

		if numeroElegido == 0 { //Condicion de salida del bucle
			break //Para que salga del bucle for
		}
	}
	fmt.Println("Has insertado 0")
}

/*
	Al igual que con el bucle while, el bucle Do While tampoco existe como tal en Go.
	Tenemos que volver a jugar con el bucle for. El bucle doWhile es una variante del while, es un bucle indefinido, donde la única diferencia es que en wl bucle while
	precisábamos de una condición para entrar al bucle, en este caso se ejecutará mínimo una vez y luego comprueba la condición.

	Teniendo en cuenta nuestro ejemplo, nuestro bucle se ejecuta mínimo una vez(te pregunta por el numero una vez mínimo) y luego si insertas 0
	te saca del bucle, si no te sigue preguntando. Con el while  necesitabamos saber que el número no era cero para que entrara en el bucle.
*/
