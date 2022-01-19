package main

import "fmt"

func add(a, b *int) int {
	fmt.Println(*a)
	fmt.Println(*b)
	*a *= 2
	*b *= 3
	fmt.Println(*a)
	fmt.Println(*b)
	return *a + *b
}

func main() {
	x, y := 1, 2
	fmt.Println(x)
	fmt.Println(y)
	z := add(&x, &y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
}
