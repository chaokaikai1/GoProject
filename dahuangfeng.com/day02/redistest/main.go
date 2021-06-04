package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func main() {
	//第一种连接方式
	con, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer con.Close()
	////第二种连接方式
	//c2,err:=redis.DialURL("redis://127.0.0.1:6379")
	//if err!=nil{
	//	panic(err)
	//}
	//defer c2.Close()
	con.Do("SET", "uname", "hello")
	uname, _ := con.Do("GET", "uname")
	if uname != nil {
		sUname := string(uname.([]byte))
		fmt.Printf("sUname is %v\n", sUname)
	}
	con.Do("SET", "uname", "helloword", "EX", "1")
	time.Sleep(5000)
	uname, _ = con.Do("GET", "uname")
	fmt.Printf("uname is %v\n", string(uname.([]byte)))

}
