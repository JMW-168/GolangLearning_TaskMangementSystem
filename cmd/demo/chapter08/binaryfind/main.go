package main

import "fmt"

// 二分查找的函數
func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	binaryfind(&arr, 0, len(arr)-1, 100)

}

func binaryfind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		binaryfind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		binaryfind(arr, middle+1, rightIndex, findVal)
	} else {
		//找到了
		fmt.Printf("找到了，下標為 %v \n", middle)
	}

}
