package main

import "fmt"

func main() {

	arr := [5]int{24, 69, 80, 57, 13}

	bubbleSort(&arr)

	fmt.Println("main arr= ", arr)

}

func bubbleSort(arr *[5]int) {

	fmt.Println("排序前 arr=", (*arr))
	temp := 0

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {

				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}
