package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	// 3.渲染模板
	u1 := User{ // Name首字母必须是大写
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{ // Name首字母可以大写也可以是小写
		"name":   "小王子",
		"gender": "男",
		"age":    18,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	err = t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}
}
