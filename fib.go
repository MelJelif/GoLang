package main

import (
	"fmt"
)

var rowfib = map[int]int{0: 0, 1: 1}

func fibfunc(n int) int {
	_, found := rowfib[n]
	if found != true {
		rowfib[n] = fibfunc(n-1) + fibfunc(n-2)
	}
	return rowfib[n]
}

func main() {

	var n int
	fmt.Println("Введите искомый пордковый номер числа: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	fmt.Println(fibfunc(n))
	fmt.Println(rowfib)
}
