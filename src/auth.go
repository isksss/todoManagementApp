package main

import (
		// "database/sql"
		"log"
)

func createUser(user Users) (err error){
	statement := "INSERT INTO users(name, email, password) VALUES(?, ?, ?);"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}
	log.Printf("insert")
	defer stmt.Close()
	log.Printf("insert")
	var id int
	err = stmt.QueryRow(user.name, user.email, user.password).Scan(&id)
	log.Printf("insert")
	// log.Printf(err)

	return
}