package main

import "fmt"

func main() {
	// goto - метка для выхода из цикла
	// Метка дожна обязательно использоваться
	// Желательно не использовать goto для большого кода
	i := 0
LOOP:
	if i > 50 {
		goto END
	}
	fmt.Println(i)
	i++
	goto LOOP
END:
	fmt.Println("Выход из цикла")
}
