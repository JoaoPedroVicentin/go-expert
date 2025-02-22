package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	cursos := Cursos{{"Go", 40}, {"HTML", 35}, {"Js", 55}}
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}

}
