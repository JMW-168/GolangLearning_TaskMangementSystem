package main

import (
	"fmt"
)

func main() {
	/*
		//切片的基本使用
		var intArr [5]int = [...]int{1, 22, 33, 66, 99}

		slice := intArr[1:3]

		fmt.Println("intArr=", intArr)
		fmt.Println("slice 的元素的是", slice)
		fmt.Println("slice 的元素個數是", len(slice))
		fmt.Println("slice 的容量", cap(slice))
	*/

	/*
		//切片的 make 應用
		var slice []float64 = make([]float64, 5, 10)
		slice[1] = 10
		slice[3] = 20

		fmt.Println(slice)
		fmt.Println("slice 的 size=", len(slice))
		fmt.Println("slice 的 cap=", cap(slice))

		//小結：
		//1.通過make 方式創建的切片可以指定切片的大小和容量
		//2.如果沒有給切片的各元素賦值，那麼就會使用默認值
		//3.通過make方式創建的切片，是由make維護，對外不可見
	*/

	//切片的第三種應用

	var strSlice []string = []string{"Ray", "Shelly", "Mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice size=", len(strSlice))
	fmt.Println("strSlice size=", cap(strSlice))

}
