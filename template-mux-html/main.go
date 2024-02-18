package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Titulo   string
	Profesor string
}

type Cursos []Curso

func main() {
	cursos := Cursos{{Titulo: "Curso de Go", Profesor: "Gustavo Henrique"}, {Titulo: "Curso de Python", Profesor: "Gustavo Henrique"}}

	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, cursos)

	if err != nil {
		panic(err)
	}
}
