package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	postForm2()
}

// 以Do的方式发送body为键值对的post请求
func postForm2() {
	uri := "http://127.0.0.1:8888"
	resource := "/test"
	data := urlValues()
	u, _ := url.ParseRequestURI(uri)
	u.Path = resource
	urlStr := u.String()
	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//http返回的response的body必须close,否则就会有内存泄露
	defer func() {
		res.Body.Close()
		fmt.Println("finish")
	}()
	//读取body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	fmt.Println(string(body))
}

func urlValues() url.Values {
	//方式1
	data1 := url.Values{"name": {"TiMi"}, "id": {"123"}}
	fmt.Println(data1)
	//方式2
	data2 := url.Values{}
	data2.Set("name", "TiMi")
	data2.Set("id", "123")
	fmt.Println(data2)
	//方式3
	data3 := make(url.Values)
	data3["name"] = []string{"TiMi"}
	data3["id"] = []string{"123"}
	fmt.Println(data3)
	/*
	   map[id:[123] name:[TiMi]]
	   map[id:[123] name:[TiMi]]
	   map[id:[123] name:[TiMi]]
	*/
	return data1
}
