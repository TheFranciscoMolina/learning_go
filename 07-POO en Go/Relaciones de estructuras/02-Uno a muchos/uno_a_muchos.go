package main

import "fmt"

type Curso struct {
	nombre string
	videos []Videos
}

type Videos struct {
	titulo string
	curso  Curso
}

func main() {

	video1 := Videos{titulo: "00-Introduccion"}
	video2 := Videos{titulo: "01-Instalacion"}

	curso := Curso{
		nombre: "APrender GO",
		videos: []Videos{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}

}
