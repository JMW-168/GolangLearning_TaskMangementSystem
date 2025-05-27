package main

import (
	"fmt"
)

func test01(arr [3]int) {
	arr[0] = 88
}

func test02(arr *[3]int) {
	(*arr)[0] = 88 //重要
}

func main() {
	// arr := [3]int{11, 22, 33}
	// test01(arr)
	// fmt.Println("main arr =", arr)
	arr := [3]int{11, 22, 33}
	test02(&arr)
	fmt.Println("main arr =", arr)
}
