package paquete

import "fmt"

/*
Aquí vamos a ver el concepto de privacidad.
Denominamos privacidad a lo que queremos y lo que no queremos que se vea fuera del paquete. Normalmente, los campos
de las estructuras suelen ser privados para que código externo al paquete no lo modifique.

Tenemos dos modificadores de visibilidad: Público y Privado.
-Público: Se define poniendo el nombre de la variable/método en mayúsculas y se puede acceder desde fuera del paquete.
-Privado: Se define poniendo el nombre de la variable/método en minúsculas y no se puede acceder desde fuera del paquete.

Cuando definimos alguna variable/campo/estructura privados solemos hacer métodos para poder acceder a ellos y usarlos.
*/

type EstructuraPublica struct {
	Id   uint
	Name string
}

type EstructuraPrivada struct {
	id   uint
	name string
}

// Como los campos de estructuraPrivada son privados, para crear la estructura definimos lo que se conoce en POO como Constructor.

func NewEstructuraPrivada(id uint, name string) *EstructuraPrivada {
	return &EstructuraPrivada{id, name}
}

//Aquí vemos la utilidad de los Setter y Getter que al hacerlos públicos nos permiten manejar los campos privados.

func (e *EstructuraPrivada) SetId(id uint) {
	e.id = id //Aquí al estar en el mismo paquete si podemos acceder directamente aunque id sea privado.
}

func (e *EstructuraPrivada) SetName(name string) {
	e.name = name
}

func (e *EstructuraPrivada) GetId() uint {
	return e.id
}

func (e *EstructuraPrivada) GetName() string {
	return e.name
}

func (e *EstructuraPrivada) String() string {
	return fmt.Sprintf("EstructuraPrivada{ id: %d, name: %s}", e.id, e.name)
}
func (e EstructuraPublica) String() string {
	return fmt.Sprintf("EstructuraPublica{ id: %d, name: %s}", e.Id, e.Name)
}
