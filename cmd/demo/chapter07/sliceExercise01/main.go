package main

import "fmt"

func main() {

	//費波那契數列練習

	fbnSlice := fbn(10)
	fmt.Println("fbnSlice=", fbnSlice)

}

func fbn(n int) []uint64 {
	//聲明一個切片
	fbnSlice := make([]uint64, n)

	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}

	return fbnSlice

}
