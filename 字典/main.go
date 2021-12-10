package main

import "fmt"

var tastMap map[string]int

func main(){
	tastMap = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	fmt.Println(tastMap)
	tastMap["four"] = 4  // 增
	fmt.Println(tastMap)
	delete(tastMap, "four")  // 删
	fmt.Println(tastMap)
	tastMap["one"] = 11  // 改
	fmt.Println(tastMap)

	fmt.Println(tastMap["one"]) // 查

	//遍历
	for k,v := range tastMap{
		fmt.Println(k,v)
	}


}