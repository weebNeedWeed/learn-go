package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER = "postgres"
	DB_PASS = "mypassword"
	DB_NAME = "sales_db"
)

func main() {
	dbInfo := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", DB_USER, DB_NAME, DB_PASS)
	db, err := sql.Open("postgres", dbInfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertedId int
	err = db.QueryRow(`INSERT INTO product_type (name) VALUES($1) RETURNING id`, "helloworld").Scan(&lastInsertedId)
	checkErr(err)
	fmt.Println(" Last inserted id = ", lastInsertedId)

	fmt.Println("# Updating")
	stmt, err := db.Prepare("UPDATE product_type SET name=$1 WHERE id=$2")
	checkErr(err)

	res, err := stmt.Exec("HEllo stmt", lastInsertedId)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(" Row Affected = ", affected)

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT id, name FROM product_type")
	checkErr(err)

	for i := 1; rows.Next(); i++ {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		checkErr(err)
		fmt.Printf(" Row %d: %d %s\n", i, id, name)
	}

	fmt.Println("# Deleting")
	stmt, err = db.Prepare("DELETE FROM product_type WHERE id=$1")
	checkErr(err)

	res, err = stmt.Exec(lastInsertedId)
	checkErr(err)
	affected, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(" Row Affected = ", affected)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
