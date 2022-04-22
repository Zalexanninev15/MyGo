package main

import (
	"fmt"
)

func main() {
	go s("Go!") // Выполненние в отдельном потоке

	// Взаимодествие с помощью канала
	ch := make(chan int)   // Создание канала для взаимодествия с потоком и тип данных
	go ss("Hello Go!", ch) // Запуск в отдельном потоке
	fmt.Println(<-ch)      // Получение данных из потока <-, Запись данных ->

	ch1 := make(chan string) // Создание канала для взаимодествия с потоком и тип данных
	ch2 := make(chan int)
	go sss("Hello Go Go!", ch1, ch2) // Запуск в отдельном потоке
	fmt.Println(<-ch1, <-ch2)        // Получение данных из потока <-, Запись данных ->

	// Закрытие канала
	// Исправление ошибки: all goroutines are asleep - deadlock!
	ch3 := make(chan int)
	go say("Let's Go!", ch3)
	for a := range ch3 {
		fmt.Println(a)
	}
}

func s(greet string) {
	fmt.Println(greet)
}

func ss(greet string, ch chan int) {
	fmt.Println(greet)
	ch <- 7
}

func sss(greet string, ch1 chan string, ch2 chan int) {
	fmt.Println(greet)
	ch1 <- "I make this!"
	ch2 <- 77
}

func say(great string, ch3 chan int) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		ch3 <- i
	}
	// Закрытие канала
	close(ch3)
}
