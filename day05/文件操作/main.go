package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 使用Read方法读取数据
	// var tmp = make([]byte,1024)
	// n, err := file.Read(tmp)
	// if err == io.EOF{
	// 	fmt.Println("文件读完了")
	// 	return
	// }
	// if err != nil {
	// 	fmt.Println("read file failed, err:", err)
	// 	return
	// }
	// fmt.Printf("读取了%d字节数据\n", n)
	// fmt.Println(string(tmp[:n]))

	// 循环读取
	var content []byte
	var tmp = make([]byte, 64)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
