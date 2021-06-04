package Models

import "time"

type Account struct {
	Id       int `gorm:"primary_key"`
	UserName string
	Age      int
	Tel      string
	AddTime  time.Time
	IsDel    int
}
