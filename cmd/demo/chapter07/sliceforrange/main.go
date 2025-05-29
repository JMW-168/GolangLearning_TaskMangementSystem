package main

import "fmt"

func main() {

	/*
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

		//用append內置函數，可以切片進行動態追加
		var slice3 []int = []int{100, 200, 300}
		fmt.Println("slice3", slice3)

		//如果要擴展具體的元素，使用append
		slice3 = append(slice3, 400, 500, 600)
		fmt.Println("slice3", slice3)
	*/

	//切片的拷貝操作
	//copy 內只能放切片
	//slice4, slice5 數據空間獨立
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)

}
