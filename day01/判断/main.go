package main

import "fmt"


// if 条件判断
func f1(){
	// score := 89
	if score := 59; score > 90 {  // 缩小作用域，减少代码冲突
		fmt.Println("优")
	} else if score == 89 {
		fmt.Println("B")		
	} else {
		fmt.Println("C")
	}
}
// for 循环
func f2(){
	// 标准for循环
	// for i := 0;i <=10;i++ {
	// 	fmt.Println(i)
	// }
	// 第二种循环
	// i := 0
	// for ;i <=10;i++ {
	// 	fmt.Println(i)
	// }
	// 第三种
	i := 0
	for i <=10 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	j := 0
	for {
		if j > 12{
			break  // 跳出循环
		}
		fmt.Println(j)
		j++
	}

	// for range

	// swich 语句
	finger := 3
	switch finger{
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("xxxx")

	}


}

func f3(){
	for i:=1;i<10;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%dx%d=%d\t",j,i,j*i)
		}
		fmt.Println()
	}
}
func main(){
	f1()
	f2()
	f3()
}