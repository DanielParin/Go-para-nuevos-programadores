package main

import "fmt"

/*
El main es el punto básico de ejecucuón del programa, es decir, siempre se buscará la función main para ejecutarse primero.
El programa se ejecutará de manera secuencial, es decir las instrucciones se leen de arriba hacia abajo.
*/
func main() {
	fmt.Println("Comentario con salto de linea.") //Nos permite mostrar por consola un String, es decir, una cadena de texto.Además nos hace un salto de línea.
	fmt.Print("Comentario sin salto de linea.")
	fmt.Println(" Sigo en la misma linea. Y luego hago un salto de linea.")

	/*
		Una variable es un espacio que almacenara un tipo de dato. Según el tipo de dato tendrá unas carácteristicas y unas operaciones asociadas a ellos,
		(no es lo mismo una suma de dos numeros enteros que  una suma de caracteres).
		Para definir una variable hay dos formas:
	*/
	var numero int  // Declarando el tipo -> var nombreVariable tipo
	texto := "Juan" // Declarando su valor -> nombreVariable := valor

	numero = 45       //Le podemos asignar valor a  las variables nuevas
	texto = "Alberto" //Las cadenas de texto o Strings van entre comillas dobles.

	fmt.Printf("Hola soy %s y tengo %d años.\n\n", texto, numero)
	/*
		El Printf nos sirve para formatear la cadena. Podemos poner que tipo de dato poner y a que variable corresponde.
		En nuestro caso nos encontramos con %s que al ser la primera, corresponderá con la primera variable que pongamos, en este caso el nombre.
		Aqui los tipos de datos más usados:

		%d -> Para numeros enteros.
		%f -> Para números de punto flotante.
		%s -> Para cadenas.
		%c -> Para caracteres.
		%t -> Para booleanos.
		%v -> Para valores en general.
		%p -> Para punteros.
	*/

	//Ahora vamos a leer los datos por pantalla:
	var nombre string
	var apellido string
	var edad int

	fmt.Print("Inserte nombre y apellido: ")
	fmt.Scanln(&nombre, &apellido)
	/*
		Aqui decimos que las dos primeras cadenas o elementos que encuentre en la consola sean el nombre y apellido.
		Si está vacía no da error pero no coge la variable, y si es de otro tipo que no sea string te lo transforma a
		string, por ejemplo un numero.
	*/
	fmt.Printf("Hola %s %s \n\n", nombre, apellido) //Por último lo imprimimos.

	fmt.Print("Inserte nombre y edad: ")
	fmt.Scanf("%s %d", &nombre, &edad)
	/*
		Aquí lo que decimos es que los elementos que leemos sean de un tipo, en este caso un string y un int.
		Por eso en este caso si no ponemos nada o ponemos otro tipo de dato que  no sea un entero cuando lo que espera es el entero
		nos pone 0 y no pasa como con el Scanln que te lo transforma.
	*/
	fmt.Printf("%s tiene %d años.", nombre, edad)

}
