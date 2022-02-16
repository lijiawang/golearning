package main

import "fmt"

//类型定义
type NewInt int

//类型别名
/*
类型别名是Go1.9版本添加的新功能。

类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。
type TypeAlias = Type
*/
type MyInt = int



func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
	/*
	结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
	b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型,也就是int类型的别名
	*/
}
