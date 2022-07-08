package main

// how to determine a linked list without close circuit
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// go connect mysql

func main() {
	connectDB()
}

func connectDB() {
	// to connect
	dsn := "root:Zhaoxiuzhu#1@tcp(127.0.0.1:3306)/sakila"

	db, err := sql.Open("mysql", dsn) // open will not verify username and password
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}

	fmt.Println("connect succesful")
}

var db *sql.DB

func query() {

}

func insert() {

}
