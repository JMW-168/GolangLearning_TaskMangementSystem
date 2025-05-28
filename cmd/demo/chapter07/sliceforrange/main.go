package main

import "fmt"

func main() {

	//使用常規方式遍歷
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	//使用for-range 方式來遍歷切片
	for i, v := range slice {
		fmt.Printf("slice[%v]=%v\n", i, v)
	}

}
