package main

import (
	"fmt"
)

// panic/recover
func main() {
	defer func() {
		r := recover() // 捕获异常
		fmt.Println(r)  // 报错信息
		fmt.Println("异常了哈哈", r)
	}()
	var ljw map[string]int
	// panic("奔溃")  # 跑出异常
	ljw["aaa"] = 1
	fmt.Println("aaaa")
}
