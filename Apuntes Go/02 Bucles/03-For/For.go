package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserte numero para sacar su tabla de multiplicar: ")
	fmt.Scanln(&numero)

	for i := 1; i <= 10; i++ {

		var resultado = numero * i

		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}

/*
	El bucle for es el único bucle que está en Go y también en otros lenguajes.
	El bucle for es un bucle definido, es decir, sabemos cuantas veces se va a ejecutar (en nuestro caso 10 veces). Este bucle se caracteriza
	por un "contador" que se suele llamar "i" y es el que nos lleva las iteraciones(veces que ejecutamos el código) de nuestro bucle for.

	Para usar el bucle precisamos de 3 parámetros(separados por punto y coma):
		-Desde cuando se inicia: Aquí solemos definir el valor inicial del contador. (i:=1)
		-Hasta cuando se ejecuta: Condición de salida del bucle(i>=10), donde ya podemos saber las iteraciones que tendrá el bucle
		-Mientras hace: Expresión de actualización, es decir, que hará cada vez que se complete una iteración, (i++)

	variable++ =>  variable = variable + 1
	variable-- => variable = variable - 1


	Por tanto la ejecución de nuestro código sería la siguiente:
		- 1ª iteración:
			El contador "i" entra valiendo 1, como es menor o igual que 10 se ejecuta y se multiplica por el número que hayamos insertado.
			Luego imprimimos y como ya hemos terminado el código se vuelve a ejecutar.
		- 2ª iteración:
			El contador "i" pasa a valer 2(por el i++), como es menor o igual que 10 se ejecuta y se multiplica por el número que hayamos insertado.
			Luego imprimimos y como ya hemos terminado el código se vuelve a ejecutar.

		..............

		-11ª iteración:
			El contador "i" pasa a valer 11 y como ya no es menor o igual que 10 sale del bucle (no lo ejecuta).
*/
