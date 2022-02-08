package main

import (
	"fmt"
	"os/user"
	"os"
)

func main() {

	u, err := user.Current() // 当前登录的用户
	fmt.Println("错误信息err: ", err)
	fmt.Println("用户ID:Uid: ", u.Uid)
	fmt.Println("用户名:Username: ", u.Username)
	fmt.Println("用户组ID:Gid: ", u.Gid)
	fmt.Println("用户组名称:Name: ", u.Name)
	fmt.Println("用户家目录:HomeDir: ", u.HomeDir)
	gids,err := u.GroupIds()
	if err != nil{
		fmt.Println("aaa")
	}
	fmt.Println(gids)
	fmt.Println(os.Getwd())  //查看当前路径
	f, errMsg1 := os.Create("test.txt")   // 创建文件
	if errMsg1 != nil {
		fmt.Println("文件创建失败", errMsg1)
		return
	} else {
		fmt.Println(f.Name(), "文件创建成功")
	}



}
