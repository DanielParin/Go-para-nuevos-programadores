package main

import "fmt"

/*
	Lo primero que vamos a hacer es, como ya hemos visto antes que las estructuras son un tipo de dato, vamos a usarlas para definir un campo de otra estructura.
*/

type Campo struct {
	nombre string
	edad   uint
}

type Estructura struct {
	id      uint
	persona Campo // persona es del tipo Campo definido antes
}

/*
	Metodos de estructura:

	Uso similar a la programación orientada a objetos, estamos "restringiendo" el uso de una función/método a una Estructura.
	Para definirlo: func (<nomVariable> <nomEstructura>) <nomMetodo>(<argumento>) <valorDevuelto> .
	No tiene por qué tener todos los elementos mencionados arriba, pero si necesitamos definir la parte entre func y <nomMetodo>, ya que
	es lo que lo asocia a la estructura.
	En este caso vamos a hacer un toString que sirve para imprimir por pantalla la estructura de la forma que queramos.
*/

func (estructura Estructura) ToString() {
	//Para ver como se llama, en vez de devolver vamos a hacer que imprima. (No es la finalidad del ToString() su funcionalidad es el otro método.
	fmt.Printf("ID: %d ; Nombre: %s  Edad: %d\n", estructura.id, estructura.persona.nombre, estructura.persona.edad)
}

// Como el ToString es un método ampliamente usado, podemos usar String() y lo podremos imprimir, sin llamarlo directamente.
func (estructura Estructura) String() string {
	//Usamos Sprintf ya que nos permite formatear una cadena que nos devuelve pero no imprime.
	return fmt.Sprintf("ID: %d ; Nombre: %s  Edad: %d\n", estructura.id, estructura.persona.nombre, estructura.persona.edad)
}

//Ahora vamos a ver en que influye el uso de punteros con estructuras.
//Los Setter y Getter también son métodos conocidos que se usan para asignar y recoger los valores de un campo.

func (estructura *Estructura) SetId(numero uint) {
	//Aqui recogemos la estructura con puntero.
	estructura.id = numero
}

func (estructura Estructura) CambiarId(numero uint) {
	//No usamos puntero.
	estructura.id = numero
}

func main() {

	//Al dar valores simplemente tenemos en cuenta que en vez de ser un valor simple es una estructura.
	estructura1 := Estructura{1, Campo{"Marcio", 15}}
	fmt.Print(estructura1)

	estructura2 := Estructura{2, Campo{"Juana", 22}}
	estructura2.ToString() // Llamamos al método con <nomVariableEstructura>.<nomMetodo>

	//ESTRUCTURA-PUNTEROS

	estructura3 := Estructura{33, Campo{"Daniel", 18}}
	estructura3.ToString()
	estructura4 := Estructura{44, Campo{"Claudia", 32}}
	estructura4.ToString()

	estructura3.SetId(3)
	estructura4.CambiarId(4)
	estructura3.ToString() //Vemos que al pasar el puntero en el método cambia el valor del ID.
	estructura4.ToString() //Al no pasarlo por puntero, el cambio se queda en el scope de la función.

	estructuraNormal := Estructura{10, Campo{"Josebas", 30}}
	estructuraNormal2 := estructuraNormal

	estructuraNormal.id = 14

	estructuraNormal.ToString()
	estructuraNormal2.ToString()

	estructuraPuntero := &Estructura{11, Campo{"Marga", 63}}
	estructuraPuntero2 := estructuraPuntero

	estructuraPuntero.id = 15 //Al haber definido la estructura como &Estructura{} creamos un puntero por lo que
	//si creamos otra estructura a partir de esa los cambios se verán reflejados.
	//En el caso anterior simplemente creamos una copia, es decir, una variable independiente.
	estructuraPuntero.ToString()
	estructuraPuntero2.ToString()
}
