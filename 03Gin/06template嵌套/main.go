package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数kua-模板自定义函数
	// 要么只有一个返回值, 要么有两个返回值, 第二个返回值必须是error类型
	k := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	// 解析模板
	t := template.New("f.tmpl") // 创建一个名字是f的模板对象, 名字一定要与定义模板的名字能对应上
	// 告诉模板引擎, 我现在多了一个自定义函数kua
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	t, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}
}
