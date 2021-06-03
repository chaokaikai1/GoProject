package main

import (
	Models2 "GoProject/daxigua.com/gostudy/Models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	//queryOne(db)
	//queryMulti(db)
	//insertData(db)
	updateData(db)
}

//查询单行
func queryOne(db *sql.DB) {
	stu := Models2.Student{}
	selSql := `select id,name,age,addr,isdel from student where id=?`
	row := db.QueryRow(selSql, 1)
	if err := row.Scan(&stu.Id, &stu.Name, &stu.Age, &stu.Addr, &stu.IsDel); err != nil { //不scan 会导致连接不释放
		panic(err)
		return
	}
	fmt.Printf("stu %#v\n", stu)
}

//查询多行
func queryMulti(db *sql.DB) {
	stu := new(Models2.Student)
	stuList := make([]Models2.Student, 0)
	selSql := `select id,name,age,addr,isdel from student where isdel=?`
	rows, err := db.Query(selSql, 0)
	defer func() {
		if rows != nil {
			rows.Close() // 可以关闭掉未scan 连接一直占用
		}
	}()
	if err != nil {
		panic(err)
		return
	}
	for rows.Next() {
		rows.Scan(&stu.Id, &stu.Name, &stu.Age, &stu.Addr, &stu.IsDel) //不scan 连接不释放
		//fmt.Println(stu)
		stuList = append(stuList, *stu)
	}
	fmt.Printf("stuList：%#v\n", stuList)
}

//插入数据
func insertData(db *sql.DB) {
	insertSql := `INSERT INTO student (name,age,addr,isdel)VALUES(?,?,?,?)`
	result, _ := db.Exec(insertSql, "ee", 55, "广州", 0) //其实到这就算执行结束了
	pkId, _ := result.LastInsertId()
	fmt.Println(pkId)
}

//更新和删除其实是一样的
func updateData(db *sql.DB) int {
	sql := `UPDATE student SET name=?,age=? WHERE id =? `
	result, _ := db.Exec(sql, "dd", 44, 4)
	count, _ := result.RowsAffected()
	fmt.Println(count)
	return int(count)
}
