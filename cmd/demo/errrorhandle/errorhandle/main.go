package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	test02()
	fmt.Println("main()下面的code")

}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}

	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res", res)

}

func readConf(name string) (err error) {
	if name == "config.ini" {
		//讀取
		return nil
	} else {
		//返回一個自定義錯誤
		return errors.New("讀取文件錯誤~")
	}
}

func test02() {

	err := readConf("config2.ini")
	if err != nil {
		panic(err)
	}

	fmt.Println("test02 繼續執行")

}
