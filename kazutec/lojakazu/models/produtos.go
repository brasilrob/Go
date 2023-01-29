package models

import "github.com/kazutec/lojakazu/db"

type Produto struct {
	Nome       stringls
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {

	db := conectaComBancoDeDados()

	selectDeTodosProdutos, err := db.Query("Select * From produtos")
	if err != nil{
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}
	
	for selectDeTodosProdutos.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil{
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer  db.Close()

	return produtos
}