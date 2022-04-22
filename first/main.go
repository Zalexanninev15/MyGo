package main // Иекущий файл с кодом

import (
	"fmt"
	"math"
) // Импорт для использования функций

func main() {
	// Присваивание - var [var] [type]
	var age int8 = 23
	fmt.Println(age)
	// Краткий способ
	age1 := 27
	fmt.Println(age1)

	// Получение ввода с клавиатуры
	var name string
	fmt.Println("Как тебя зовут?")
	fmt.Scan(&name) // & - указатель
	fmt.Println("Привет, " + name + "!")
	fmt.Println("Сколько тебе лет?")
	fmt.Scan(&age)
	fmt.Println("Мне уже " + fmt.Sprint(age)) // fmt.Sprint - конвертация в строку
	// конвертирование в другие типы [тип в который надо сконвертировать]([переменна для конвертации])

	// Цикл
	for i := 0; i <= 10; i++ { // цикл от 0 до 10 включительно
		println(i)
	}
	for i := 0; i <= 10; i += 2 { // цикл от 0 до 10 включительно, увеличение i с шагом 2
		println(i)
	}
	for i := 10; i >= 0; i-- { // цикл от 10 до 0 включительно
		println(i)
	}
	// Объявление переменной i заранее
	var i1 int = 0
	for ; i1 <= 50; i1 += 10 { // цикл от 0 до 50, увеличение i с шагом 10
		fmt.Println(i1)
	}
	i1 = 0
	for i1 <= 50 { // Цикл также можно создать и так
		fmt.Println(i1)
		i1 += 10
	}

	// Отложенный вызов функции
	defer fmt.Println("в конце выполнения приложения")
	fmt.Println("Функция вызывается отложено")

	// Switch
	var h string = "A"
	switch h {
	case "A":
		fmt.Println("Is 'A'")
	case "B":
		fmt.Println("Is 'B'")
	case "C":
		fmt.Println("Is 'C'")
	default:
		fmt.Println("Is ?")
	}
	var l int8 = 3
	switch {
	case l > 2:
		fmt.Println("Число больше 2")
		fallthrough // Продолжить выаолнение для других case при совпадении
	case l < 10:
		fmt.Println("Число меньше 10")
	}

	// Округление числа
	var a float64 = 58.27845
	result := math.Ceil(a) // Округление к большему значению
	fmt.Println(result)
	result = math.Floor(a) // Округление к меньшему значению
	fmt.Println(result)
	result = math.Round(a) // Обычное округление
	fmt.Println(result)

	// Форматирование
	money := 1002.492
	bool_value := false
	// Для int - %d
	// Для string - %s
	// Для float - %f
	// Для bool - %t
	fmt.Printf("My age is %d. My name is %s. I have %f dollars. Value is %t", age, name, money, bool_value)
}
