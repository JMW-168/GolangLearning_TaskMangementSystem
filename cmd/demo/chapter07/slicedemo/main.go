package main

import (
	"fmt"
)

func main() {

	//切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}

	slice := intArr[1:3]

	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素的是", slice)
	fmt.Println("slice 的元素個數是", len(slice))
	fmt.Println("slice 的容量", cap(slice))
}
