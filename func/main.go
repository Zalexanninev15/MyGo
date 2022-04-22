package main

import "fmt"

func function(fun func(int) int) { // Вызов функции из функции
	fmt.Println(fun(25))
}

func main() {
	fmt.Println(first_func(1, 6))
	sm, sb, mu, di := second_func(6, 4) // Получение нескольких значений из функции
	fmt.Println(sm, sb, mu, di)
	sm, sb, mu, di = second_func(2, 8) // Получение нескольких значений из функции
	fmt.Println(sm, sb, mu, di)
	a := func() { // Создание функции внутри другой функции (main)
		fmt.Println("Love it!")
	}
	a()
	square := func(x int) int {
		return x * x
	}
	function(square)
	second_second_func("Hello")()
}

func first_func(a int, b int) int { // Возвращаемое значение
	result := a + b
	return result
}

func second_func(a int, b int) (int, int, int, int) { // Возврат нескольких значений
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	return sum, sub, mul, div
}

func second_func1(a int, b int) (sum int, sub int, mul int, div int) { // Возврат нескольких значений (2 способ)
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	return
}

func second_second_func(x string) func() { // Вызов функции из функции в функции
	return func() {
		fmt.Println(x)
	}
}
