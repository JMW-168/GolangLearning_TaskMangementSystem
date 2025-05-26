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
}
