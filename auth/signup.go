package auth

import (
	"log"

	"github.com/isksss/todoManagementApp/db"
)

func (user Users) CreateUser() (err error) {
	statement := "INSERT INTO users(id, name, email, password) VALUES(?, ?, ?, ?);"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	// 実行
	res, err := stmt.Exec(user.Id, user.Name, user.Email, user.Password)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 結果の取得
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("CREATE_USER_ID: ", lastInsertID)
	return
}
