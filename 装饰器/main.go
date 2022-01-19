package main
import (
	"fmt"
	"time"
)

// 为函数类型设置别名提高代码可读性
type MultiPlyFunc func(int,int) int

func multiply(a,b int) int{
	return a * b
}

func execTime(f MultiPlyFunc) MultiPlyFunc {
    return func(a, b int) int {
        start := time.Now() // 起始时间
        c := f(a, b)  // 执行乘法运算函数
        end := time.Since(start) // 函数执行完毕耗时
        fmt.Printf("--- 执行耗时: %v ---\n", end)
        return c  // 返回计算结果
    }
}
func main() {
    a := 2
    b := 8
    // 通过修饰器调用乘法函数，返回的是一个匿名函数
    decorator := execTime(multiply)
    // 执行修饰器返回函数
    c := decorator(a, b)
    fmt.Printf("%d x %d = %d\n", a, b, c)
}