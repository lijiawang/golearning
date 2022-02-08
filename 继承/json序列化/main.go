package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   string 
	Name string
	Sex  string
}

type Class struct {
	Name     string
	Count    int
	Students []*Student
}

var raqJson = `
{"Name":"101","Count":0,"Students":[{"Id":"10","Name":"stu0","Sex":"man"},{"Id":"11","Name":"stu1","Sex":"man"},{"Id":"12","Name":"stu2","Sex":"man"},{"Id":"13","Name":"stu3","Sex":"man"},{"Id":"14","Name":"stu4","Sex":"man"},{"Id":"15","Name":"stu5","Sex":"man"},{"Id":"16","Name":"stu6","Sex":"man"},{"Id":"17","Name":"stu7","Sex":"man"},{"Id":"18","Name":"stu8","Sex":"man"},{"Id":"19","Name":"stu9","Sex":"man"}]}
`
func main() {
	fmt.Println("hahaha")

	c := &Class{
		Name:  "101",
		Count: 0,
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			Sex:  "man",
			Id:   fmt.Sprintf("1%d", i),
		}
		c.Students = append(c.Students, stu)
	}

	data,err := json.Marshal(c)
	if err != nil{
		fmt.Println("json marsha failed")
		return
	}
	fmt.Printf("json:%s\n",string(data))

	// 反序列化
	fmt.Println("反序列化")
	var c1 *Class = &Class{}
	err = json.Unmarshal([]byte(raqJson),c1)

	if err != nil{
		fmt.Println("unmarhsal failed")
		return
	}
	fmt.Printf("c1:%#v\n",c1 )
	
	for _,v := range c1.Students{
		fmt.Printf("stu:%#v\n",v.Name)
	}
}
