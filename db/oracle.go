package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
)

func commitTest() {
	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("oci8", "cms_master/cms_master@192.168.8.72:1522/orcl?autocommit=true&charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	_, err = tx.Exec("update MT_NODE set DIMENSIONS=39.038853 where NODE_ID=25")
	if err != nil {
		fmt.Println(err)
	}

	_, err = tx.Exec("update MT_NODE set DIMENSIONS=39.075739 where NODE_ID=26")
	if err != nil {
		fmt.Println(err)
	}

	_, err = tx.Exec("update MT_NODE set DIMENSIONS=39.075739 where NODE_ID=29")
	if err != nil {
		fmt.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func main() {
	commitTest()
}
