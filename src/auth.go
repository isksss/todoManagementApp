package main

// "database/sql"
import (
	"log"
	// "database/sql"
)

// CreateUser
func (user Users) createUser() (err error) {
	statement := "INSERT INTO users(name, email, password) VALUES(?, ?, ?);"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Fatal("db error.")
	}
	defer stmt.Close()

	// 実行
	res, err := stmt.Exec(user.name, user.email, user.password)
	if err != nil {
		log.Fatalln(err)
	}

	// 結果の取得
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CREATE_USER_ID: ", lastInsertID)
	return
}

// login
// func (user Users) login_auth() (err error) {
// 	// statement := "SELECT password FROM users WHERE name = ? OR email = ?"
// }
