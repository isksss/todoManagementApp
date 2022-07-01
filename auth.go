package main

// "database/sql"
import (
	"fmt"
	"log"
	// "database/sql"
)

// CreateUser
func (user Users) createUser() (err error) {
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

// login
func (user Users) loginAuth() (bl bool) {
	statement := "SELECT password FROM users WHERE name = ? OR email = ?"

	stmt, err := Db.Prepare(statement)
	// showError(err)
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := stmt.Query(user.name, user.email)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var password string
	for rows.Next() {
		rows.Scan(&password)
		// log.Printf("DB GET PASSWORD: %s\n", )
		fmt.Println(&password)
		log.Printf("USER GET PASSWORD: %s\n", user.password)
		if user.password == password {
			return true
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	stmt.Close()

	return false
}
