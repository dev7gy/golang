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
		fmt.Println(colorToString(Yellow))
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
	/*
		case문에 여러 값 사용 가능
		조건문도 사용 가능
	*/
	switch c := n % 2; c {
	case 0:
		fmt.Println("This is even number", n)
		// break // break를 사용하든 안하든 go언어에서는 break가 기본임
	case 1:
		fmt.Println("This is odd number", n)
		fallthrough
	default:
		fmt.Println("This is default. by fallthrough")
	}
}

/*
const 열거값과 Switch
*/
type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "undefined"
	}
}
