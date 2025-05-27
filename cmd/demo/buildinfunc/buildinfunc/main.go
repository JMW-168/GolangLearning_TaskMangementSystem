package main

import (
	"fmt"
)

func main() {

	num1 := 100
	fmt.Printf("num 的類型 %T\n,num1 的值是 %v\n, num1 的地址為 %v\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2 的類型 %T\n,num2 的值是 %v\n, num2 的地址為 %v\n", num2, num2, &num2)

}
