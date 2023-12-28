package main

import "fmt"

func main() {
	// 기본 및 중첩 for문
	for i := 0; i < 3; i++ {
		fmt.Println("i", i)
		// 초기문 생략
		j := 0
		for ; j < 2; j++ {
			fmt.Println("j", j)
		}
	}

	// 레이블을 활용한 중첩 for문 전부 break
	// 후처리 생략
BreakForFor:
	for k := 0; k < 5; {
		fmt.Println("k", k)
		k++
		// 조건문만 있는 경우
		l := 0
		for l < 5 {
			fmt.Println("l", l)
			if l == 3 {
				break BreakForFor // 레이블처리
			}
			l++
		}
	}

	// continue break
	for m := 0; m < 5; m++ {
		if m%2 == 0 {
			continue
		}
		fmt.Println(m)
	}

	/*
	 무한 루프
	 for { }를 사용하는 것을 권장
	*/
	n := 0
	for true {
		if n == 3 {
			break
		}
		fmt.Println("infinite ", n)
		n++
	}

	n = 0
	for {
		if n == 2 {
			break
		}
		fmt.Println("infinite2 ", n)
		n++
	}

}
