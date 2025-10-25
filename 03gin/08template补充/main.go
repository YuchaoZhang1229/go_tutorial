package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("index.tmpl")
	if err != nil {
		fmt.Println("parse template err:", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.Execute(w, name)
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 解析模板之前定义一个自定义的函数safe
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("xss.tmpl")
	if err != nil {
		fmt.Println("parse template err:", err)
		return
	}
	// 渲染模板
	str1 := `<script>alert('嘿嘿嘿')</script>`
	str2 := "<a href='http://liwenzhou.com'>liwenzhou的博客</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}
}
