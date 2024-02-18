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

	t := template.Must(template.New("CursoTemplate").Parse("Este {{.Titulo}} é ministrado pelo professor {{.Profesor}}"))
	err := t.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}
