package main

import (
 "GoProject/dahuangfeng.com/Models"
 "encoding/json"
 "fmt"
)

func main() {
 //stu:=Models.Student{
 //  ID: 1,
 //  Name: "aa",
 //  Gender: "男",
 //}
 //data,_:=json.Marshal(stu)
 //jsonResult:=string(data)
 //fmt.Println(jsonResult)
 //stu2:=Models.Student{}
 //json.Unmarshal(data,&stu2)
 // fmt.Printf("stu2:%#v",stu2)
 stu:=Models.Student{}
 jsonStr:=`{"id":1,"gender":"男","name":"aa"}`
 json.Unmarshal([]byte(jsonStr),&stu)
 fmt.Printf("stu2:%#v",stu)


}
