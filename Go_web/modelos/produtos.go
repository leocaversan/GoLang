package modelos

import (
	_ "Go_web/db"
	db2 "Go_web/db"
	"fmt"
)

type Produto struct {
	Id         int
	Nome       string
	Marcas     string
	Preco      float64
	Quantidade int64
}

func BuscaTodosProdutos() []Produto {
	db := db2.Connect_my_sql()

	select_all_products, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for select_all_products.Next() {
		var id int
		var nome, marca string
		var preco float64
		var quantidade int64

		err = select_all_products.Scan(&id, &nome, &marca, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Marcas = marca
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos
}

func CriarNovosProdutos(nome string, descricao string, preco float64, quantidade int) {
	db := db2.Connect_my_sql()

	insereInfoBancoDados, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) values (?, ?,?,?)")

	if err != nil {
		fmt.Println("erro ao conect")
		panic(err.Error())
	}
	insereInfoBancoDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func UpdateBancoNewValues(nome string, descricao string, preco float64, quantidade int, id int) {
	db := db2.Connect_my_sql()

	insereInfoBancoDados, err := db.Prepare("UPDATE produtos SET nome=?, descricao=?, preco=?, quantidade=? WHERE ID=?")

	if err != nil {
		fmt.Println("erro ao conect")
		panic(err.Error())
	}
	insereInfoBancoDados.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}

func DeletarProduto(nome string) {
	db := db2.Connect_my_sql()

	deletaInfoBancoDados, err := db.Prepare("DELETE FROM produtos WHERE nome = ?")

	if err != nil {
		fmt.Println("erro ao conect")
		panic(err.Error())
	}
	deletaInfoBancoDados.Exec(nome)
	defer db.Close()
}

func DeleteProdutoById(id string) {
	db := db2.Connect_my_sql()

	deletaInfoBancoDados, err := db.Prepare("DELETE FROM produtos WHERE id = ?")

	if err != nil {
		fmt.Println("erro ao conect")
		panic(err.Error())
	}
	deletaInfoBancoDados.Exec(id)
	defer db.Close()
}

func UpdateProdutoById(idProduto string) Produto {
	db := db2.Connect_my_sql()

	produtoBanco, err := db.Query("SELECT * FROM produtos WHERE id = ? ", idProduto)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			fmt.Println("ERRO ao Ler produtos")
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Marcas = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = int64(quantidade)
	}
	defer db.Close()
	return produtoParaAtualizar
}
