package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Sort(ar []int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
func main() {
	fmt.Println("Введите числа которые нужно отсортировать через запятую: ")
	var srs string
	var sls = []int{}
	_, err := fmt.Scanln(&srs)
	if err != nil {
		panic(err)
	}
	nums := strings.Split(srs, ",")
	for _, num2 := range nums {
		num, _ := strconv.Atoi(num2)
		//код написан с использовнием копирования массива, но в данном случае считаю это не критичным
		sls = append(sls, num)
		//fmt.Println(len(sls), cap(sls), unsafe.Pointer(&sls[0])) оставил строку для проверки
	}
	Sort(sls)
	fmt.Println(sls)
}
