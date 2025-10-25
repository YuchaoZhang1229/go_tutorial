


当我们遇到比如说像index.tmpl和home.tmpl只有部分不一样的时候, 我们可以使用模板继承 bolck 语法
1. 在./templates/base.tmpl定义一个根模板, 将可替换的内容改成
```html
{{block "content" .}}{{end}}
```
2. 对应./templates/index2.tmpl和./templates/home2.tmpl代码如下
```html
// index2
{{/*继承根模板 加个.把base模板的数据继承过来*/}}
{{template "base.tmpl" .}}
{{/*重新定义块模板*/}}
{{define "content"}}
    <h1>这是Index2页面</h1>
    <p>Hello {{ . }}</p>
{{end}}
```


