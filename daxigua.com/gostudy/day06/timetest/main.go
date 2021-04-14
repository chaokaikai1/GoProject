package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//add
	now2 := now.Add(2 * time.Hour) //加两小时
	now2 = now.Add(time.Hour)      //加1小时
	fmt.Println(now2)

	//时间戳
	fmt.Println(now.Unix())               //时间to时间戳
	fmt.Println(time.Unix(now.Unix(), 0)) //将时间戳转换为时间

	//时间格式化
	timeStr := now.Format("2006-01-02 15:04:05")      //2021-04-14 23:26:26
	timeStr2 := now.Format("2006-01-02 03:04:05")     //2021-04-14 11:26:26
	timeStr3 := now.Format("2006-01-02 15:04:05 PM")  //2021-04-14 23:26:26 PM
	timeStr4 := now.Format("2006-01-02 03:04:05 PM")  //2021-04-14 11:26:26 PM
	timeStr5 := now.Format("2006-01-02 15:04:05.000") //2021-04-14 23:27:59.318
	fmt.Println(timeStr)
	fmt.Println(timeStr2)
	fmt.Println(timeStr3)
	fmt.Println(timeStr4)
	fmt.Println(timeStr5)

}
