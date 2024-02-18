package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Titulo   string
	Profesor string
}

type Cursos []Curso

func main() {
	cursos := Cursos{{Titulo: "Curso de Go", Profesor: "Gustavo Henrique"}, {Titulo: "Curso de Python", Profesor: "Gustavo Henrique"}}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, cursos)

		if err != nil {
			panic(err)
		}

	},
	)

	http.ListenAndServe(":8080", mux)

}
