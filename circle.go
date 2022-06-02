package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	b := math.Pi

	fmt.Print("Введите площадь круга: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Диаметр окружности %.2f \n", math.Sqrt(a/b)*2)

}
