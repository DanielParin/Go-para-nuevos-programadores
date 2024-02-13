package main

import "fmt"

var variable1 = "Soy una variable de nivel 1."

func main() {

	variable2 := "Soy una variable de nivel 2."
	condicion := true

	if condicion { //Condicion para que siempre entre.
		variable3 := "Soy una variable de tipo 3."
		fmt.Println(&variable1)
		fmt.Println(&variable2)
		fmt.Println(&variable3)
	}

	fmt.Println(&variable1)
	fmt.Println(&variable2)
	//fmt.Println(&variable3) //Descomenta el primer comentario de esta línea y mira el error.

}

/*
	Definimos Scope de una variable al rango/alcance de uso que tiene esta. Una variable no puede ser utilizada en un entorno fuera de las llaves en donde se definen,
	si no se usan otras funcionalidades que veremos más adelante.

	Por ejemplo el Scope de la variable2 sería  todo lo que esté dentro de las llaves ({}) del main, si se intentara usar fuera de su entorno nos daría un error,
	como nos pasa con la variable3. Si se intenta usar la variable3 fuera de las llaves del if, donde está definida, nos salta error.
*/
