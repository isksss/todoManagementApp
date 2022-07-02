package auth

import (
	"log"
)

type Users struct {
	id       int
	name     string
	email    string
	password string
}

func (user *Users) createUser() (err error) {
	statement := "INSERT INTO users(name, email, password) VALUES(?, ?, ?);"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	// 実行
	res, err := stmt.Exec(user.name, user.email, user.password)
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
