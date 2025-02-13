package main

/*
	Podríamos definir los paquetes como una carpeta con código dentro (para entendernos).
	El paquete fmt por ejemplo es un paquete que nos agrega características para utilizar la entrada y salida formateada.
	Al ser un paquete propio de Go lo podemos importar directamente, pero como en este caso para importar nuestro paquete
	tenemos que seguir la estructura de carpetas desde el módulo definido en el fichero go.mod

	Si quisiéramos usar algún paquete externo a nosotros y a Go, debemos hacer un: go get <nombrePaquete> y este se añadirá
	al go.mod
*/
import (
	"fmt"
	"main/Paquetes/paquete"
)

func main() {

	estructura := paquete.EstructuraPublica{1, "Dan"}
	fmt.Printf("El Id de la estructura publica es: %d\n", estructura.Id)
	fmt.Printf("El nombre de la estructura pública es: %s\n", estructura.Name)
	estructura.Name = "Daniel"
	//Al ser campos públicos podemos acceder a ellos directamente.

	fmt.Println(estructura)

	estructura2 := paquete.NewEstructuraPrivada(2, "Mara")
	fmt.Printf("El Id de la estructura privada es: %d\n", estructura2.GetId())
	fmt.Printf("El nombre de la estructura privada es: %s\n", estructura2.GetName())
	estructura2.SetName("Dibella")
	//Al ser campos privados tenemos que acceder a ellos a través de métodos públicos.

	fmt.Println(estructura2)

}
