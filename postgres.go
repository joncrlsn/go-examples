//
// This is an example of accessing Postgres using the standard database/sql API.
// Also included is a method for dynamically converting each row to a string array.
//
package main

import "fmt"

//import "reflect"
import "time"
import "os"
import "database/sql"
import _ "github.com/lib/pq"

const isoFormat = "2006-01-02T15:04:05.000-0700"

func main() {
	db, err := sql.Open("postgres", "user=dbuser password=supersecret dbname=dev-cpc sslmode=disable")
	panicIfError(err)
	readUsersStatically(db)
	readUsersDynamically(db)
}

func readUsersStatically(db *sql.DB) {
	var userId int
	var username string

	maxUserId := 21
	rows, err := db.Query("SELECT user_id, username FROM t_user WHERE user_id < $1", maxUserId)
	panicIfError(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userId, &username)
		panicIfError(err)
		fmt.Println(userId, username)
	}

	err = rows.Err()
	panicIfError(err)
}

/* Dynamically convert each row in the result set into a string array */
func readUsersDynamically(db *sql.DB) {
	fmt.Println("--- dynamic column retrieval ---")

	maxUserId := 4
	rows, err := db.Query("SELECT user_id, username, active, creation_date, intro_email_date FROM t_user WHERE user_id < $1", maxUserId)
	panicIfError(err)
	defer rows.Close()

	// http://stackoverflow.com/questions/23507531/is-golangs-sql-package-incapable-of-ad-hoc-exploratory-queries
	columnNames, err := rows.Columns()
	panicIfError(err)

	vals := make([]interface{}, len(columnNames))
	valPointers := make([]interface{}, len(columnNames))
	for i := 0; i < len(columnNames); i++ {
		valPointers[i] = &vals[i]
	}
	for rows.Next() {
		err = rows.Scan(valPointers...)
		panicIfError(err)
		cells := make([]string, len(columnNames))
		// Convert each cell to a SQL-valid string representation
		for i, valPtr := range vals {
			//fmt.Println(reflect.TypeOf(valPtr))
			switch valueType := valPtr.(type) {
			case nil:
				cells[i] = "null"
			case []uint8:
				cells[i] = "'" + string(valPtr.([]byte)) + "'"
				//fmt.Println("Value is a []uint8 (string)", cells[i])
			case string:
				cells[i] = "'" + valPtr.(string) + "'"
				//fmt.Println("Value is a String", cells[i])
			case int64:
				cells[i] = string(valPtr.(int64))
				//fmt.Println("Value is an int64", cells[i])
			case bool:
				cells[i] = fmt.Sprintf("%t", valPtr)
				//fmt.Println("Value is a bool", cells[i])
			case time.Time:
				cells[i] = "'" + valPtr.(time.Time).Format(isoFormat) + "'"
				//fmt.Println("Value is a time.Time", cells[i])
			case fmt.Stringer:
				cells[i] = fmt.Sprintf("%v", valPtr)
				//fmt.Println("Value is an fmt.Stringer", cells[i])
			default:
				cells[i] = fmt.Sprintf("%v", valPtr)
				fmt.Printf("Column %s is an unhandled type: %v", columnNames[i], valueType)
			}
		}
		fmt.Println("Cells:", cells)
	}
}

func panicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		//panic(err)
	}
}
