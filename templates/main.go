package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Titulo   string
	Profesor string
}

func main() {
	curso := Curso{Titulo: "Curso de Go", Profesor: "Gustavo Henrique"}

	template := template.New("CursoTemplate")
	template, _ = template.Parse("Este {{.Titulo}} é ministrado pelo professor {{.Profesor}}")
	err := template.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
