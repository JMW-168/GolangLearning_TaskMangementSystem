package main

import (
	"fmt"
)

/*
func main() {
	a := 20
	b := 15
	res1, res2 := calc.GetSumAndSub(a, b)
	fmt.Printf("a 和 b 的總和是 %d \na 和 b 的差是 %d\n", res1, res2)

	//如果希望忽略某個返回值，使用 _ 符號表示占位忽略

	_, res3 := calc.GetSumAndSub(1, 2)
	fmt.Println("res3", res3)

	//遞迴調用實作
	recursiontest(7)
	recursiontest2(7)

	//遞迴調用練習，費波那契數列實作
	//給你一個n，求出他的費波那契數列
	febo := feboNumber(10)
	fmt.Println("febo=", febo)

}

func recursiontest(n int) {
	if n > 2 {
		n--
		recursiontest(n)
	}
	fmt.Println("n=", n)
}

func recursiontest2(n int) {
	if n > 2 {
		n--
		recursiontest2(n)
	} else {
		fmt.Println("n2=", n)
	}

}

func feboNumber(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return feboNumber(n-1) + feboNumber(n-2)
	}
}
*/

//猴子吃桃問題
//猴子有一堆桃子，每天都吃一半多一個，第10還沒吃的時候，發現剩下一顆，請問他原本有幾顆

//n為天數
//規律 ： peach(n) = (peach(n+1)＋1) * 2

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("數量不對")
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}

}

func main() {
	fmt.Println("猴哥第1天的桃子數量為:", peach(1))
}
