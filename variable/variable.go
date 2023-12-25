package main

import "fmt"

func main() {

	fmt.Println("The default values for each type are as follows.")
	var integer8 int8
	fmt.Printf("int8: %d\n", integer8)

	var unsignedInteger16 uint16
	fmt.Printf("uint16: %d\n", unsignedInteger16)

	var floatNum float32
	fmt.Printf("float32 %f\n", floatNum)

	var isBool bool
	fmt.Printf("bool %t\n", isBool)

	var str string
	fmt.Printf("string %s\n", str)

	typeChange()
}

func typeChange() {
	var a float32 = 3.5
	fmt.Println(a)
	b := int(a) // b == 3
	fmt.Println(b)

	var c int16 = 3456
	fmt.Println(c)

	var d int8 = int8(c)
	fmt.Println(d)

}
