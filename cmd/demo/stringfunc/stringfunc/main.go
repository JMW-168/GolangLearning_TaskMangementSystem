package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		str := "hello北七"
		//統計字符串的長度，中文一字三個字節，英數一字一個字節
		fmt.Println("str lens=", len(str))

		str2 := "hello西安"
		//字符串遍歷，因為中文三字節，會被分開有問題，使用 r:=[]rune(str)
		r := []rune(str2)

		for i := 0; i < len(r); i++ {
			fmt.Printf("字符=%c\n", r[i])
		}
	`	*/
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr = %v\n", strArr)
}
