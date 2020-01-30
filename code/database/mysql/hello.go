package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "reflect"

type Timestamp []uint8

func main() {

	db, err := sql.Open("mysql", "root:@/information_schema")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// db.QueryRow("SELECT CURRENT_TIMESTAMP FROM DUAL WHERE 1 = ?", 1)

	//
	// TEST: 1
	//
	if true {
		rows, err := db.Query("SELECT CURRENT_TIMESTAMP FROM DUAL WHERE 1 = ?", 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		for rows.Next() {
			var value interface{}
			if err = rows.Scan(&value); err != nil {
				fmt.Println(err)
			} else {
				// How can I Scan() TIMESTAMP ??
				fmt.Println(reflect.TypeOf(value))
			}
		}
	}

	//
	// TEST: 2
	//
	if false {
		statement, err := db.Prepare("SELECT CURRENT_TIMESTAMP WHERE 1 = ?")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer statement.Close()
		statement.QueryRow(1)
	}

	fmt.Println("Ok.")
}
