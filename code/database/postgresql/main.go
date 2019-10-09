package main

import "fmt"
import "log"
import "database/sql"
import _ "reflect"
import _ "github.com/lib/pq"

func main() {

    fmt.Println("[TRACE] ### START ###")

	db, err := sql.Open(
		"postgres",
		"host=127.0.0.1 port=5432 user=user1 password=password dbname=testdb sslmode=disable")

    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

	rows, err := db.Query("SELECT CURRENT_TIMESTAMP WHERE 1 = $1", 1)

	if err != nil {
		fmt.Println(err)
		return		
	}

	for rows.Next() {
		var value interface{}
		if err = rows.Scan(&value); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("[TRACE] %v\n", value)
	}

    fmt.Println("[TRACE] --- END ---")
}
