package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL ")
	serverName := "localhost:3306"
	user := "root"
	password := ""

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	//db, err := sql.Open("mysql", "root:password @tcp(127.0.0.1:3306)/org")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO test VALUES ( 4, 'Bhuniya' )")


	if err != nil {
		panic(err.Error())
	}
	
	defer insert.Close()
	// Execute the query
	results, err := db.Query("SELECT id, name FROM test")
	if err != nil {
		panic(err.Error()) 
	}

	for results.Next() {
		var tag test
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) 
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}

	var tag test
	// Execute the query
	err = db.QueryRow("SELECT id, name FROM test where id = ?", 2).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) 
	}

	log.Println(tag.ID)
	log.Println(tag.Name)

}
