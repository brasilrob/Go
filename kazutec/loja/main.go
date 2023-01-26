package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8089", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 100},
		{"Tenis", "Confortável", 80, 80},
		{"Fone", "Muito Bom", 59, 2},
		{"Produto Novo", "Muito Legal", 1.99, 2},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
