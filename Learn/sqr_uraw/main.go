package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Квадратное уравнение")

	var a float64
	var b float64
	var c float64

	fmt.Println("Введи a: ")
	fmt.Scan(&a)
	fmt.Println("Введи b: ")
	fmt.Scan(&b)
	fmt.Println("Введи c: ")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уравнение имеет 2 корня\nD = " + fmt.Sprint(D))
		fmt.Println("X1 = " + fmt.Sprint(x1) + "\nX2 = " + fmt.Sprint(x2))
	} else if D == 0 {
		var x1 float64

		x1 = (-b) / (2 * a)

		fmt.Println("Ваше уравнение имеет 1 корень\nD = 0")
		fmt.Println("X1 = " + fmt.Sprint(x1))
	} else if D < 0 {
		fmt.Println("Ваше уравнение не имеет корней!\nD = " + fmt.Sprint(D))
	}

	fmt.Scan() // Чтобы консоль не закрывалась
}
