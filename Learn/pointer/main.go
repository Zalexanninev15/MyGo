package main

import "fmt"

func change(str *string) { // Функция с оператором разменования
	*str = "Changed"
}

func main() {
	// Указатель - адрес переменной в памяти
	a := 5
	pointer := &a
	fmt.Println(pointer)  // Имя ячейки памяти в которой хранится переменная 'a'
	fmt.Println(*pointer) // Значение переменной
	// * - оператор разименования

	// Изменение значения с помощью переменной и указателя с оператором разименования
	// Значения сохраняется в той же ячейке памяти, а не записывается в другую
	s := "Value"
	fmt.Println(s)
	change(&s)
	fmt.Println(s)
	// Или так:
	// var pointer *string = &s
	// change(pointer)
}
