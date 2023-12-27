package main

import "fmt"

func main() {

	a, b := 13, 2
	fmt.Println(a, b)
	fmt.Println("Add: ", a, "+", b)
	fmt.Println(Add(a, b))
	fmt.Println("Divide: ", a, "/", b)
	fmt.Println(DivideA(a, b))
	fmt.Println(DivideB(a, b))

	RecursiveRange(3)
}

func Add(a int, b int) int {
	return a + b
}

// 멀티 반환 함수
func DivideA(a int, b int) (bool, int, int) {
	var quotient int
	var remainder int
	var isSuccess bool = false
	if b == 0 || b > a {
		return isSuccess, quotient, remainder
	}
	isSuccess = true
	quotient = a / b
	remainder = a % b

	return isSuccess, quotient, remainder
}

// 변수명 지정 반환 함수
func DivideB(a int, b int) (isSuccess bool, quotient int, remainder int) {
	// 변수명을 지정해서 반환하기
	isSuccess = false
	quotient = 0
	remainder = 0
	if b == 0 || b > a {
		return
	}

	isSuccess = true
	quotient = a / b
	remainder = a % b

	return
}

// 재귀 함수
func RecursiveRange(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	RecursiveRange(n - 1)
	fmt.Println("After", n)
}
