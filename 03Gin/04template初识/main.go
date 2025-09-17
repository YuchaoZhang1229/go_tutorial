package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	// 3.渲染模板
	err = t.Execute(w, "沙河小王子")
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
