package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil { // nil - пустота (null)
		fmt.Println("Не могу прочитать файл - ", err)
		os.Exit(1) // Выход с ошидкой
	}
	defer file.Close() // Закрытие соединения с файлом
	file.WriteString("Text to file\n1234")
	file_data, _ := os.ReadFile(file.Name())
	fmt.Println(file_data)         // Вывод байтовый срез
	fmt.Println(string(file_data)) // Вывод текста
}
