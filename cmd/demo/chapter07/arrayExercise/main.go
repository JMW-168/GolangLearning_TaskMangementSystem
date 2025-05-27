package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		//列出a-z
		var mychars [26]byte
		for i := 0; i < 26; i++ {
			mychars[i] = 'A' + byte(i)
		}

		for i := 0; i < 26; i++ {
			fmt.Printf("%c", mychars[i])
		}
	*/

	/*
		//請求出一個陣列的最大值並得到其對應的下標（位置）
		var intArr [5]int = [...]int{1, -1, 9, 90, 11}
		var maxitem = intArr[0]
		var maxitemIndex = 0

		for i := 0; i < len(intArr); i++ {
			if intArr[i] > maxitem {
				maxitem = intArr[i]
				maxitemIndex = i
			}
		}
		fmt.Printf("陣列最大值為%v, maxitemIndex=%v \n", maxitem, maxitemIndex)
	*/

	//求出一個數組的平均
	var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	fmt.Printf("sum=%v 平均值＝%v\n", sum, float64(sum)/float64(len(intArr2)))

	var intArr3 [5]int

	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)
}
