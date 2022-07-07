package auth

import (
	"log"
	"net/http"

	"github.com/isksss/todoManagementApp/db"
	cr "github.com/isksss/todoManagementApp/crypto"
)

type Users struct {
	Id       int
	Name     string
	Email    string
	Password string
}

// login
func (user Users) CheckLogin() (bl bool) {
	bl = false
	statement := "SELECT password FROM users WHERE name = ? OR email = ?"

	stmt, err := db.Db.Prepare(statement)
	// showError(err)
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := stmt.Query(user.Name, user.Email)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var password string
	for rows.Next() {
		rows.Scan(&password)
		if user.Password == password {
			bl = true
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	stmt.Close()

	return
}

func (user Users) MakeSession(w http.ResponseWriter) {
	
	sess := cr.NewSessionKey()
	
	statement := "INSERT INTO sessions(session_id, user_id) VALUE (?, ?);"
	

}
