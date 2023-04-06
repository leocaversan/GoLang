package main

import (
	_ "Go_web/modelos"
	"Go_web/routes"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
