package main

import "fmt"

func main() {
	/*
		var score [5]float64

		for i := 0; i < len(score); i++ {
			fmt.Printf("請輸入第%d個元素的值\n", i+1)
			fmt.Scanln(&score[i])
		}

		for i := 0; i < len(score); i++ {
			fmt.Printf("score[%d]=%v\n", i, score[i])
		}
	*/
	//四種初始化陣列的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr01=", numArr02)

	var numArr03 = [...]int{1, 2, 3}
	fmt.Println("numArr01=", numArr03)

	var numArr04 = [...]int{1: 800, 0: 700, 2: 100}
	fmt.Println("numArr01=", numArr04)

	numArr05 := [...]string{1: "Tom", 2: "Terry", 3: "Ray"}
	fmt.Println("numArr01=", numArr05)

}
