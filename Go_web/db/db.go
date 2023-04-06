package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect_my_sql() *sql.DB {
	conexao := "root:@tcp(127.0.0.1:3306)/DBO_GOLANG"
	db, err := sql.Open("mysql", conexao)

	if err != nil {
		fmt.Println("Erro ao conectar")
		fmt.Println(err)
		panic(err.Error())
	}
	return db

}
