package main

// how to determine a linked list without close circuit
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// go connect mysql
func main() {
	err := connectDB()
	if err != nil {
		fmt.Printf("open DB failed, err:%v\n", err)
		return
	}
	fmt.Println("connect succesful")

	// queryOneRow(201)
	// queryMultiple(3)
	// insert()
	// update("ting", 201)
	// delete(201)
	// prepareInset()
	transactionDemo()
}

func connectDB() (err error) {
	// to connect
	dsn := "root:Zhaoxiuzhu#1@tcp(127.0.0.1:3306)/sakila"

	db, err = sql.Open("mysql", dsn) // open will not verify username and password
	if err != nil {
		return
	}

	err = db.Ping()

	if err != nil {
		return
	}

	// set max connection fo db
	db.SetMaxOpenConns(10)

	return
}

type actor struct {
	actor_id    int
	first_name  string
	last_name   string
	last_update string
}

// CRUD create, read, update and delete
//  queryOneRow
func queryOneRow(id int) {

	sqlStr := `select actor_id, first_name, last_name,  last_update from actor where actor_id=?;`

	// get result from row
	var a1 actor
	err := db.QueryRow(sqlStr, id).Scan(&a1.actor_id, &a1.first_name, &a1.last_name, &a1.last_update) // must scan the rowObj // query data from connection pool
	if err != nil {
		fmt.Println("queryrow error:", err)
		return
	}

	fmt.Printf("a1%#v\n", a1)
}

func queryMultiple(n int) {

	sqlStr := `select actor_id, first_name, last_name,  last_update from actor where actor_id>?;`

	// get result from row
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("query '%s' error:%s", sqlStr, err)
		return
	}

	// must close the rows
	defer rows.Close()

	// loop rows
	for rows.Next() {
		var b1 actor
		err := rows.Scan(&b1.actor_id, &b1.first_name, &b1.last_name, &b1.last_update)
		if err != nil {
			fmt.Println("queryrow error:", err)
			return
		}
		fmt.Printf("a1%#v\n", b1)
	}
}

// insert data
func insert() {
	// 1. write sql
	sqlStr := `insert into actor (first_name, last_name,  last_update) values ("xiang", "wang", "2022-07-09 04:34:33")`

	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("inser failed, err:%v\n", err)
		return
	}

	/// get the last id

	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}

	fmt.Println("id:", id)
}

// update data

func update(newFirstName string, id int) {
	sqlStr := `update actor set first_name=? where actor_id=?`
	ret, err := db.Exec(sqlStr, newFirstName, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("update %d sucessful", n)
}

func delete(id int) {
	sqlStr := `delete from actor where actor_id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("update %d sucessful", n)
}

// prepare
func prepareInset() {
	sqlStr := `insert into actor (first_name, last_name,  last_update) values (?, ?, "2022-07-09 04:34:33")`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	m := map[string]string{
		"user1": "xiang",
		"user2": "ting",
		"user3": "x1",
		"user4": "t1",
	}

	for k, v := range m {
		stmt.Exec(k, v)
	}
}

// go transaction
func transactionDemo() (_, err error) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return nil, err
	}

	// exec multiple sql operation
	sqlStr1 := `update actor set first_name="xxx" where actor_id=202`
	sqlStr2 := `update actor set last_name="YYY" where actor_id=202`

	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("first sql query %s, err:%v\n", sqlStr1, err)
		return nil, err
	}

	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("first sql query %s, err:%v\n", sqlStr2, err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
		return nil, err
	}

	// if both with success, the commit
	fmt.Print("transction succeed")
	return
}
