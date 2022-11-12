package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var route string = os.Getenv("DB_ROUTE")
	var database string = os.Getenv("DB_DATABASE")

	var dbCredential string = fmt.Sprintln("%s:%s@tcp(%s)/%s", user, password, route, database)

	db, dbError := sql.Open("mysql", dbCredential)

	if dbError != nil {
		log.Fatal(dbError.Error())
	}

	pingError := db.Ping()
	if pingError != nil {
		log.Fatal(pingError.Error())
		fmt.Println("Error")
	}

	db.Query("SHOW TABLES;")

	_, errorCreate := db.Exec("CREATE TABLE livros(" +
		"id INT NOT NULL AUTO_INCREMENT," +
		"author VARCHAR(50) NOT NULL," +
		"title VARCHAR(50) NOT NULL," +
		"PRIMARY KEY(id)" +
		");")

	if errorCreate != nil {
		log.Fatal(errorCreate.Error())
	}

	_, erroInsert := db.Exec("INSERT INTO livros (author, title) VALUES " +
		"('Harry Potter','J.K')," +
		"('Clear Code','Uncle Bob')," +
		"('Frankenstein','Mary Shelley');")

	if erroInsert != nil {
		log.Fatal(erroInsert.Error())
	}
}
