package main

import "fmt"

func main() {
	//string 底層是一個byte 數組 ，因此可以進行切片處理！
	str := "Terry@WJ"
	slice := str[6:]
	fmt.Println("slice=", slice)
}
