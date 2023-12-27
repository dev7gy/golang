package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	fmt.Println(a, b)
	a, b = b, a // 복수 대비 연산자
	fmt.Println(a, b)

	a += 1 // 복합 대비 연산자
	fmt.Println(a)
	a-- // 증감 연산자
	fmt.Println(a)

	var c int8 = 30
	fmt.Println(c)
	c <<= 2
	c += 8 // int8이다보니 오버플로 발생
	fmt.Println(c)
}
