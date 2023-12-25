package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input int
	stdin := bufio.NewReader(os.Stdin) // 버퍼지우기 용도
	for i := 0; i < 10; i++ {
		n, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(n, err)
			fmt.Println("error")
			stdin.ReadString('\n')
		} else {
			fmt.Println(i, input)
		}
	}
}
