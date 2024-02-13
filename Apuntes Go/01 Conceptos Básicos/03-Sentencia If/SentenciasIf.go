package main

import "fmt"

func main() {

	numeroComparacion := 2 //Prueba a cambiarlo

	if numeroComparacion == 2 {
		fmt.Println("El numero es un 2.")
	} else {
		fmt.Println("El numero no es 2")
	}
	/*
		La sentencia if es la más básica de las sentencias. Su función es comparar dos valores, en nuestro caso, si numeroComparacion tiene el valor 2
		nos ejecutará el código que hay dentro del if.
		¿Qué pasa si no se cumple la condición? Lo primero es que no se ejecuta el código del if, y después depende de si pones la sentencia else
		o no. Si pones la sentencia else se ejecutará su código independientemente de cual sea su valor, si no la pones simplemente no hace el código
		del if.
	*/

	var valorCondicion bool
	/*
		El tipo booleano solo puede tener dos valores true o false.
		Lo que nos sirve para poder usar como comparación en un If, si es true ejecuta el if, si es false no lo hace.
	*/

	valorCondicion = true //Prueba a cambiarlo por false

	if valorCondicion {
		fmt.Println("El booleano es True!!")
	} //Si fuese false no ejecutaría el if

	//También podemos poner varias condiciones
	if valorCondicion && numeroComparacion == 2 {
		fmt.Println("Las condiciones se cumplen.")
	}
	/*
		        OPERADORES PARA CONDICIONES:

		        !=  Que sea distinto, como tal la ! delante significa no, en este caso que no sea igual
						se puede usar también con booleanos poniendolo delante de la variable ( !variable ).
		        ==  Que sea igual, no confundir con un solo igual, ya que eso no es una condición sino una asignación
		              de valor.
		        <   Que sea menor que , no incluye el propio numero.
		        >   Que sea mayor que, no incluye el propio numero.

		        <=, >=  Que sea menor igual/mayor igual que, incluye el número.

		        ENTRE CONDICIONES:

		        x && z -> si pasa x y pasa z me lo haces, con que  una no se cumpla ya no lo hace.

		        x || z -> si pasa x o pasa z me lo haces, con que se cumpla minimo una lo hace.

	*/

	if numeroComparacion == 1 {
		fmt.Print("El numero es 1.")
	} else if numeroComparacion == 2 {
		fmt.Print("El numero es 2.")
	} else if numeroComparacion == 3 {
		fmt.Print("El numero es 3.")
	} else {
		fmt.Print("El numero no está entre los 3 primeros.")
	}

	/*
		La sentencia else if se utiliza para seguir comparando despues de un if. Si ocurre me haces esto(if), si no ocurre y si pasa esto otro me haces esto(else if).
		Es similar a dentro de la ejecución del else poner un if y así sucesivamente. El else final es cuando no cumple  ninguna de las otras condiciones.
	*/
}
