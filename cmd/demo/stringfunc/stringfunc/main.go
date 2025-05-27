package main

import (
	"fmt"
	"time"
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
	/*
		//依照特定字符分割
		strArr := strings.Split("hello,world,ok", ",")
		for i := 0; i < len(strArr); i++ {
			fmt.Printf("str[%v]=%v\n", i, strArr[i])
		}
		fmt.Printf("strArr = %v\n", strArr)
	*/
	//取得時間
	//now := time.Now()
	//fmt.Printf("now=%v now Type=%T", now, now)

	//結合sleep 使用時間常量
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		//time.sleep(time.second)
		time.Sleep(time.Millisecond * 500)
		if i == 100 {
			break
		}
	}
}
