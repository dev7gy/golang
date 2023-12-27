package main

import "fmt"

const PI float32 = 3.14

// iota는 소괄호를 벗어나면 다시 초기화됨.
const (
	Red   int = iota // 0
	Blue  int = iota // 1
	Green int = iota // 2
)

func FuncConst() {
	fmt.Print(Red)
	fmt.Print(Blue)
	fmt.Print(Green)
}

const ( // 똑같은 규칙이 계속 적용된다면 타입과 iota를 생략할 수 있다.
	C1 uint = iota * 2 // 0 * 2
	C2                 // 1 * 2
	C3                 // 2 * 2
)

func FuncConst2() {
	fmt.Print(C1)
	fmt.Print(C2)
	fmt.Print(C3)
}

const NOTYPE_PI = 3.14

/*
	go 언어에서 상수는 리터럴과 같이 취급

그래서 컴파일딜 때 상수는 리터럴로 변환되어 실행 파일에 쓰임
상수 표현힉 역시 컴파이 타임에 실제 결과값 리터럴로 변환하기 때문에 상수 표현식 계싼에 CPU 자원 사용하지 않음.
*/
func main() {
	var n int = NOTYPE_PI * 100 // TYPE 업는 상수를 이용할 때는 int 곱하기 가능
	fmt.Println(n)
	fmt.Println(PI)
	FuncConst()
	fmt.Println()
	FuncConst2()
}
