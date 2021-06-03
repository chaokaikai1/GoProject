package main

import (
	"GoProject/dahuangfeng.com/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getFun() string {
	url := "http://daduhui.mangotechdemo.com.au/webapi/Address/GetModelById/2"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	result := string(data)
	//fmt.Println(result)
	return result
}

func postFun() {
	strurl := "http://daduhui.mangotechdemo.com.au/webapi/Account/Login"
	contentType := "application/x-www-form-urlencoded"
	resp, _ := http.Post(strurl, contentType, strings.NewReader("username=admin&password=123456"))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	result := string(data)
	fmt.Println(result)
}

func postFun2() {
	strurl := "http://daduhui.mangotechdemo.com.au/webapi/Account/Login"
	params := url.Values{
		"username": {"admin"},
		"password": {"123456"},
	}
	resp, _ := http.PostForm(strurl, params)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	result := string(data)
	fmt.Println(result)
}
func httpDoFun1() {
	url1 := "http://daduhui.mangotechdemo.com.au/webapi/Address/GetModelById/2"
	client := http.Client{}
	req, err := http.NewRequest("get", url1, nil)
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	result := string(data)
	fmt.Println(result)
}

func httpDoFun2() {
	url1 := "http://daduhui.mangotechdemo.com.au/webapi/Account/Login"
	client := http.Client{}
	body1 := strings.NewReader("username=admin&password=123456")
	req, _ := http.NewRequest("post", url1, body1)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	result := string(data)
	fmt.Println(result)
}

func main() {

	//postFun()
	//postFun2()
	//httpDoFun1()
	//httpDoFun2()
	result:= getFun()
	fmt.Println(result)
	modelResult:= Models.ModelResult{}
	json.Unmarshal([]byte(result),&modelResult)
	fmt.Printf("result:%#v",modelResult)


}
