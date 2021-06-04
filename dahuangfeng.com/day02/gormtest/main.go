package main

import (
	"GoProject/dahuangfeng.com/Models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func main() {
	creatModel()
}

//创建记录
func creatModel() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, _ := gorm.Open("mysql", dsn)
	defer db.Close()
	account := Models.Account{
		UserName: "cc",
		Age:      33,
		Tel:      "15011540049",
		AddTime:  time.Now(),
		IsDel:    0,
	}
	db.NewRecord(account)
	//db.Create(&account)
	account2 := new(Models.Account)
	db.Find(&account2)
	fmt.Println(*account2)

}
