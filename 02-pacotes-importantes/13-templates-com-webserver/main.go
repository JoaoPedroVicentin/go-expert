package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{{"Go", 40}, {"HTML", 35}, {"Js", 55}}
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, cursos)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)

}
