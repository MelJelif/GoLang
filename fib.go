package main

import (
	"fmt"
)

type fibMap = map[int]int

func fibfunc(rowfib fibMap, n int) int {
	_, found := rowfib[n]
	if found != true {
		rowfib[n] = fibfunc(rowfib, n-1) + fibfunc(rowfib, n-2)
	}
	return rowfib[n]
}

func main() {
	var rowfib = fibMap{0: 0, 1: 1}
	var n int
	fmt.Println("Введите искомый пордковый номер числа: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	fmt.Println(fibfunc(rowfib, n))
	fmt.Println(rowfib)
}
