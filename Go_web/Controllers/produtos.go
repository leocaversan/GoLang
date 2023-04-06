package Controllers

import (
	"Go_web/LoginQuery"
	"Go_web/modelos"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := modelos.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func PagaDeleteProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "delete_2", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		modelos.DeletarProduto(nome)
	}
	http.Redirect(w, r, "/index", 301)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	modelos.DeleteProdutoById(idProduto)
	http.Redirect(w, r, "/index", 301)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produtosupdate := modelos.UpdateProdutoById(idProduto)
	temp.ExecuteTemplate(w, "update", produtosupdate)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("ERRO ao converter valores", err)
		}
		quantidadeConvertidaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("ERRO ao converter quantidade", err)
		}
		modelos.CriarNovosProdutos(nome, descricao, precoConvertido, quantidadeConvertidaInt)
	}

	http.Redirect(w, r, "/", 301)
}

func UpdateBanco(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("ERRO ao converter valores", err)
		}
		quantidadeConvertidaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("ERRO ao converter quantidade", err)
		}
		modelos.UpdateBancoNewValues(nome, descricao, precoConvertido, quantidadeConvertidaInt, idInt)
	}

	http.Redirect(w, r, "/index", 301)
}

func TelaLogin(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "login", nil)

}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		login := r.FormValue("username")
		password := r.FormValue("password")
		_, err := LoginQuery.ConsultaLogin(login, password)
		if err != nil {
			http.Redirect(w, r, "/", 301)
			return
		}
	}
	http.Redirect(w, r, "/index", 301)

}
