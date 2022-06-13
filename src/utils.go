package main


import (
	"net/http"
	"fmt"
	"html/template"
	"crypto/sha512"
	"encoding/hex"
	"log"
	
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql" //  for MySQL

	"github.com/joho/godotenv"
	
)

/*
generateHTML:
テンプレートの解析などを行う。
*/
func generateHTML(w http.ResponseWriter, data interface{}, file ...string){
	var files []string
	for _, file := range file{
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, "layout", data)
}

func passwdEncrypt(passwd string) string {

	hashed := "" // ハッシュの初期化

	sha512 := sha512.Sum512([]byte(passwd))
	hashed = hex.EncodeToString(sha512[:])

	return hashed
}

/*
DataBase:
*/
var Db *sql.DB
func init() {
	/*
		データベース接続用のメソッド
		[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	*/
	// getenv
	loadEnv()
	db := DbLogin{
		id:       os.Getenv("DB_ID"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbname:   os.Getenv("DB_NAME"),
	}
	dbconf := db.id + ":" + db.password + "@tcp(" + db.host + ":" + db.port + ")/" + db.dbname
	var err error
	Db, err = sql.Open("mysql", dbconf)

	if err != nil {
		panic(err)
	}
}


/*
環境変数読み込み:
*/
func loadEnv() {
	/*  for devlop.
		if you deploy this, uncomment and comment.
	*/
	err := godotenv.Load("./.env.dev")
	//err := godotenv.Load("./.env")

	if err != nil {
		log.Printf("ERROR: Don't load Env")
	}else{
		log.Printf("SUCCESS: Load Env")
	}
}
