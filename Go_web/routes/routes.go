package routes

import (
	"Go_web/Controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/index", Controllers.Index)
	http.HandleFunc("/new", Controllers.New)
	http.HandleFunc("/delete_2", Controllers.PagaDeleteProduct)
	http.HandleFunc("/insert", Controllers.Insert)
	http.HandleFunc("/Delete_", Controllers.Delete)
	http.HandleFunc("/delete", Controllers.DeleteById)
	http.HandleFunc("/update", Controllers.UpdateById)
	http.HandleFunc("/updateBanco", Controllers.UpdateBanco)
	http.HandleFunc("/", Controllers.TelaLogin)
	http.HandleFunc("/logar", Controllers.Login)
}
