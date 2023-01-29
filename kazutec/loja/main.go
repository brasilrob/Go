package main

import (	
	"net/http"
	"text/template"
	
	"github.com/kazutec/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
// 	db := conectaComBancoDeDados()
// 	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8089", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	
/*  produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 100},
		{"Tenis", "Confortável", 80, 80},
		{"Fone", "Muito Bom", 59, 2},
		{"Produto Novo", "Muito Legal", 1.99, 2},
		{"Laptop", "Muito fixe", 1000.99, 2},
	} */
	todosProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
	defer db.Close()
}
