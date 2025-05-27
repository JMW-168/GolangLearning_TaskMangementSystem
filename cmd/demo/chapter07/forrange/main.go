package main

import "fmt"

func main() {

	//用 for-range 來遍歷數組
	heroes := [...]string{"宋江", "吳用", "盧俊義"}

	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v) //較推薦

		fmt.Printf("heroes[%d]=%v\n", i, heroes[i])
	}

	for _, v := range heroes {
		fmt.Printf("元素的值=%v\n", v)
	}
}
