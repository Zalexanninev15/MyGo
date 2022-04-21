package main

import (
	"fmt"
	"sort"
)

func main() {
	// Массив
	list := [3]string{"Git", "GitHub", "Gitea"}
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

	// Многомерный массив
	array := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	fmt.Println(array[1][2]) // Из 2 массива вывести 3 элемент (по-человечески)

	// Срез (массив с возможностью добавления новых элементов)
	slice := []int{3, 1, 2, 5, 7, 4}
	slice = append(slice, 0)
	sort.Ints(slice) // Сортировка элементов в срезе по увеличению
	fmt.Println(slice)
	slice1 := []string{"a", "k", "v", "c", "d", "b"}
	sort.Strings(slice1) // Сортировка в алфавитном порядке
	fmt.Println(slice1)

	// Перебор среза
	slice2 := []int{3, 4, 5, 2, 6, 9, 10}
	for i, element := range slice2 { // element == slice2[i]
		fmt.Printf("%d: %d\n", i, element) // Форматирование строки %d (int)
	}
	for _, element := range slice2 { // Перебор без i
		fmt.Printf("%d\n", element)
	}

	// Карты (для производительности, вместо среза и массива)
	var money map[string]int = map[string]int{ // Типы: [[ключ]][значение]
		"dollars": 1000, // [ключ][значение]
		"euros":   2000,
		"apples":  3,
	} // порядок не сохраняется, значения будут располагаться по возрастанию
	fmt.Println(money)
	fmt.Println(money["euros"])
	delete(money, "apples") // удаление по ключу apples
	fmt.Println(money)
	// element - значение по ключу, если он существует
	// status - bool, существование значения по ключу
	element, status := money["dollars"]
	fmt.Println(element, status)
	element1, status1 := money["ruble"]
	fmt.Println(element1, status1)
}
