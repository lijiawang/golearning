package main

import (
	"fmt"
)

type User struct{
	Username string
	Age      int
	Sex      int
	Avatarurl string
}

func main(){
	var user User
	user.Username = "lijiawang"
	user.Age=18
	user.Sex = 1
	user.Avatarurl="http://ljw.com/1"
	fmt.Println(user.Username,user.Age,user.Sex,user.Avatarurl)

	var user1 User
	fmt.Printf("%#v\n",user1)
}