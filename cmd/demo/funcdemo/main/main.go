package main

import (
	"fmt"

	"github.com/chengjinming/GolangLearning_TaskMangementSystem/cmd/demo/funcdemo/utils/calc"
)

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
