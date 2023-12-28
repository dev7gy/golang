package main

import "fmt"

func main() {
	fmt.Println("EvenOddChecker")
	for {
		var n int
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		OddEvenCheckerUsingIf(n)
		OddEvenCheckerUsingSwitch(n)
	}
}

func OddEvenCheckerUsingIf(n int) {
	if c := n; c == 0 { // 초기문이 true일 때 수행됨
		fmt.Println("This is Zero")
	} else if c%2 == 0 {
		fmt.Println("This is even number", n)
	} else {
		fmt.Println("This is odd number", n)
	}
}

func OddEvenCheckerUsingSwitch(n int) {
	var c = n % 2
	/*
		case문에 여러 값 사용 가능
		조건문도 사용 가능
	*/
	switch c {
	case 0:
		fmt.Println("This is even number", n)
	case 1:
		fmt.Println("This is odd number", n)
	}
}
