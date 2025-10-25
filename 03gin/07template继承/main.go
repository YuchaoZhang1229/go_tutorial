package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("index.tmpl")
	if err != nil {
		fmt.Println("parse template err:", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.Execute(w, name)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("home.tmpl")
	if err != nil {
		fmt.Println("parse template err:", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.Execute(w, name)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板(模板继承的方式)
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println("parse template err:", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.ExecuteTemplate(w, "index2.tmpl", name)

}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板(模板继承的方式)
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println("parse template err:", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.ExecuteTemplate(w, "home2.tmpl", name)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}
}
