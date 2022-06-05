package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Введите размеры прямоугольника через пробел: ")
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		panic(err)
	}

	fmt.Println("Площадь прямоугольника: ", a*b)

}
