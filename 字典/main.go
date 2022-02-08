package main

import (
	"fmt"
)

var tastMap map[string]int

func main() {
	tastMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(tastMap)
	tastMap["four"] = 4 // 增
	fmt.Println(tastMap)
	delete(tastMap, "four") // 删
	fmt.Println(tastMap)
	tastMap["one"] = 11 // 改
	fmt.Println(tastMap)

	fmt.Println(tastMap["one"]) // 查

	//遍历
	for k, v := range tastMap {
		fmt.Println(k, v)
	}
	k := "two1"
	v, ok := tastMap[k]
	fmt.Println(ok)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}
	var aaa [10]int //定义数组
	fmt.Println("aaaaa")
	fmt.Println(len(aaa))
	aaa[1] = 2
	fmt.Println(aaa)

	var bbb []int //定义切片
	bbb = append(bbb, 123)
	var slice []string = []string{"a", "b", "c"}
	fmt.Println(slice)

	// 定义一个数据
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// 基于数组创建切片
	q2 := months[3:6]
	fmt.Println(cap(q2))
	q2 = append(q2, "bbb")

	fmt.Println(cap(q2))

	fmt.Println(q2)

	q4 := months[10:]
	q4 = append(q4, "aaa")
	q4 = append(q4, "bbb")
	q4 = append(q4, "xxx")
	fmt.Println(len(q4), cap(q4), q4)

	mySlice1 := make([]int, 5) // 创建一个长度为5，容量为5的切片

	fmt.Println(len(mySlice1), cap(mySlice1), mySlice1)

	mySlice2 := make([]int, 5, 10) // 创建一个长度为5，容量为10的切片

	fmt.Println(len(mySlice2), cap(mySlice2), mySlice2)

	mySlice3 := []int{1, 2, 3, 4, 5} //创建并初始化切片 5 5 [1 2 3 4 5]

	fmt.Println(len(mySlice3), cap(mySlice3), mySlice3)

	mySlice4 := append(mySlice3, 6, 7, 8, 9)
	fmt.Println(len(mySlice4), cap(mySlice4), mySlice4)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	// 复制 slice1 到 slice 2  只会复制 slice1 的前3个元素到 slice2 中
	// copy(slice2, slice1)
	// fmt.Println(slice1,slice2)
	copy(slice1, slice2)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice3 = slice3[:len(slice3)-5] // 删除 slice3 尾部 5 个元素
	fmt.Println(slice3)

	slice11 := []int{1, 2, 3, 4, 5}
	slice21 := slice11[1:3]
	slice21[1] = 6
	fmt.Println("slice11:", slice11)
	fmt.Println("slice21:", slice21)
	fmt.Println("slice11:", slice11)

}
