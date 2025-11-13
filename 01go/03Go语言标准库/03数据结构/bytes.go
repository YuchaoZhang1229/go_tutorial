package main

import (
	"bytes"
	"fmt"
)

func bytesOperation() {
	fmt.Println("=== bytesOperation ===")
	// 字节切片
	data1 := []byte("Hello, World!")
	data2 := []byte("Hello")
	data3 := []byte("World")

	// 比较字节切片
	fmt.Println("data1和data2比较:", bytes.Compare(data1, data2))

	// 检查前缀和后缀
	fmt.Println("data1是否以Hello开头:", bytes.HasPrefix(data1, data2))
	fmt.Println("data1是否以World!结尾:", bytes.HasSuffix(data1, []byte("World!")))

	// 包含检查
	fmt.Println("data1是否包含World:", bytes.Contains(data1, data3))

	// 计数
	fmt.Println("l出现的次数:", bytes.Count(data1, []byte("l")))

	// 替换
	replaced := bytes.Replace(data1, []byte("World"), []byte("Golang"), -1)
	fmt.Println("替换后:", string(replaced))

	// 分割和连接
	parts := bytes.Split(data1, []byte(","))
	fmt.Printf("分割结果: %q\n", parts)

	joined := bytes.Join(parts, []byte(" - "))
	fmt.Println("连接后:", string(joined))

	// 大小写转换
	fmt.Println("大写:", string(bytes.ToUpper(data1)))
	fmt.Println("小写:", string(bytes.ToLower(data1)))
	fmt.Println()
}

func bytesBuffer() {
	fmt.Println("=== bytesBuffer ===")
	// 创建缓冲区
	var buf bytes.Buffer

	// 写入数据
	buf.WriteString("Hello")
	buf.WriteByte(' ')
	buf.Write([]byte("World!"))

	fmt.Println("缓冲区内容:", buf.String())
	fmt.Println("缓冲区长度:", buf.Len())

	// 读取数据
	firstByte, _ := buf.ReadByte()
	fmt.Printf("第一个字节: %c\n", firstByte)

	// 剩余内容
	fmt.Println("剩余内容:", buf.String())

	// 重置缓冲区
	buf.Reset()
	buf.WriteString("新的内容")
	fmt.Println("重置后:", buf.String())
}

func main() {
	bytesOperation()
	bytesBuffer()
}
