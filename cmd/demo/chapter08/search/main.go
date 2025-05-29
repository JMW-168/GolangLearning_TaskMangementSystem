package main

import "fmt"

func main() {
	//有一個數列：白眉鷹王，金毛獅王，紫衫龍王，清液蝠王

	names := [4]string{"白眉鷹王", "金毛獅王", "紫衫龍王", "清液蝠王"}
	var heroName = ""
	fmt.Println("請輸入要查找的人名")
	fmt.Scanln(&heroName)
	/*
		//順序查找：第一種方式
		for i := 0; i < len(names); i++ {
			if heroName == names[i] {
				fmt.Printf("找到%v , 下標%v  \n", heroName, i)
				break
			} else if i == (len(names) - 1) {
				fmt.Printf("沒有找到%v \n", heroName)
			}
		}
	*/

	//順序查找：第2種方式
	index := -1

	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v , 下標%v  \n", heroName, index)
	} else {
		fmt.Println("沒有找到\n", heroName)
	}
}
