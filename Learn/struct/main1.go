package main

import "fmt"

// Интерфейс
// (нет данных, есть только методы)
// Интерфес может унаследовать свои методы для нескольких структур

type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	Sum() int
	Multiply() int
	Division() float64
	Substract() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}

func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}

func (n Numbers) Division() float64 {
	return float64(n.num1 / n.num2)
}

func (n Numbers) Substract() int {
	return n.num1 - n.num2
}

func main() {
	var i NumbersInterface
	nums := Numbers{9, 8}
	i = nums
	fmt.Printf("Суммирование %d\n", i.Sum())
	fmt.Printf("Умножение %d\n", i.Multiply())
	fmt.Printf("Поделить %f\n", i.Division())
	fmt.Printf("Отнять %d", i.Substract())
}
