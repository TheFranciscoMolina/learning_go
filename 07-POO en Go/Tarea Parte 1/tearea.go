package main

import "fmt"

//Struct

type Tarea struct {
	titulo      string
	descripcion string
	competado   bool
}

//Para mostrar todas las tareas:
type ListaDeTareas struct {
	//creo un slice de tipo Tarea:
	tareas []*Tarea
}

//Agregar tareas a la lista
func (app *ListaDeTareas) agregaTarea(acc *Tarea) {
	app.tareas = append(app.tareas, acc)
}

//Eliminar tareas de esa lista:
func (app *ListaDeTareas) quitaTarea(index int) {
	app.tareas = append(app.tareas[:index], app.tareas[index+1:]...)
}

//Metodo para imprimir tareas ordenadas:
func (acc *Tarea) paraImprimir() {
	fmt.Printf("Tarea: %s\nDescripcion:%s\n, Completado: %t\n", acc.titulo, acc.descripcion, acc.competado)
}

func (acc *Tarea) darCompletado() {
	acc.competado = true
}

func main() {
	verCurso := Tarea{
		titulo:      "Curso de Go",
		descripcion: "Completar el curso de Go desde las bases hasta poder hacer tareas de back.",
		competado:   false,
	}

	sacarFotos := Tarea{
		titulo:      "Hacer foto de productos",
		descripcion: "Sacar todos a los productos que me manden.",
		competado:   true,
	}
	aprenderCSS := Tarea{
		titulo:      "Aprender CSS",
		descripcion: "Hacer curso de estilos en cascada",
		competado:   false,
	}

	/* verCurso.paraImprimir()
	sacarFotos.paraImprimir() */

	primeraLista := ListaDeTareas{}
	primeraLista.agregaTarea(&verCurso)
	primeraLista.agregaTarea(&aprenderCSS)
	primeraLista.agregaTarea(&sacarFotos)
	primeraLista.quitaTarea(1)

	// fmt.Println("La lista tiene:", primeraLista)

	for i, actividad := range primeraLista.tareas {
		fmt.Println(i, actividad.titulo)
	}

}
