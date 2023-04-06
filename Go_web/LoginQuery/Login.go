package LoginQuery

import (
	_ "Go_web/db"
	db2 "Go_web/db"
	"fmt"
)

type User struct {
	Username string
	Password string
}

func ConsultaLogin(login string, password string) (User, error) {
	db := db2.Connect_my_sql()
	defer db.Close()
	var user User
	err := db.QueryRow("SELECT login, senha FROM login WHERE login=? AND senha=?", login, password).Scan(&user.Username, &user.Password)
	fmt.Println(err)
	if err != nil {
		return user, err
		fmt.Println(err)
	}
	return user, nil

}
