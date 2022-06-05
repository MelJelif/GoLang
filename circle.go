package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("Введите площадь круга: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Диаметр окружности %.2f \n", math.Sqrt(a/math.Pi)*2)
	fmt.Printf("Длина окружности : %.2f \n", (math.Sqrt(a/math.Pi)*2)*math.Pi)
}
