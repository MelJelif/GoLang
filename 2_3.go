package main

import (
	"fmt"
	"math"
)

func main() {
	var a int

	fmt.Print("Введите трехзначное число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		panic(err)
	}
	b := math.Floor(float64(a % 1000 / 100))
	c := math.Floor(float64(a % 100 / 10))
	d := math.Floor(float64(a % 10))
	fmt.Printf("В числе %v, %v сотен, %v десятков, %v единиц\n", a, b, c, d)
}
