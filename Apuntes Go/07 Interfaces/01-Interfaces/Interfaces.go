package main

import "fmt"

/*
	Para ver que son las interfaces vamos a poner un ejemplo.
	Imaginemos que tenemos dos estructuras, un Gato y un Perro. Los dos pueden emitir sonido, pero no es el mismo sonido
	cada una tendrá el suyo propio, pero las dos emiten sonido. Por ello podemos hacer una interfaz que sea animal
	que englobe a las estructuras gato y perro donde ahí se personalice el método emitir sonido.
	En la programación orientada a objetos esto se conoce como POLIMORFISMO.

	Las interfaces solo pueden contener métodos y para que se implemente una interfaz en una estructura, la estructura
	tiene que utilizar todos los métodos de la interfaz.
*/

// Definimos la interfaz con type <nomInterfaz> interface
type Animal interface {
	EmitirSonido() string //Método que devuelve un string
}

type Perro struct{}

func (p Perro) EmitirSonido() string {
	return "Guau!!"
}

type Gato struct{}

func (g Gato) EmitirSonido() string {
	return "Miau!!"
}

func main() {

	var animal1 Animal = Perro{}
	var animal2 Animal = Gato{}

	/*
		Como ambas estructuras implementan la interfaz Animal, se puede asignar el tipo Animal a un perro, pero no
		viceversa, todos los perros son animales pero no todos los animales son perros.
	*/

	fmt.Println(animal1.EmitirSonido())
	fmt.Println(animal2.EmitirSonido())
}
