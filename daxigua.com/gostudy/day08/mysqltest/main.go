package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn %s invalid, err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")
}
