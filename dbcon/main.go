package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func OpenMysql() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", `ubox_report:vG6ke8di3bWopRP3fCb6@tcp(211.151.164.45:3307)/ubox_alipaycs?autocommit=true&charset=utf8`)
	if db != nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
	}
	return
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("Error panic: %v\n", x)
		}
	}()

	db, err := OpenMysql()
	if err != nil {
		fmt.Println("aaaa" + err.Error())
	}

	count := 0

	sql_ := `SELECT COUNT(0) FROM ub_alipaycs_cancel`
	stmt, err := db.Prepare(sql_)
	if err != nil {
		fmt.Printf("Error exec sql: %s\n", err.Error())
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRow().Scan(&count); err != nil {
		fmt.Printf("Error exec sql: %s\n", err.Error())
	}

	fmt.Println(count)
	db.Close()
}
