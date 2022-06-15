package main

// "database/sql"
import (
	"log"
	// "database/sql"
)

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
	log.Println(lastInsertID)
	return
}
