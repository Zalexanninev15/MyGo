package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
	score    []int
}

func (u *User) setName(name1 string) { // Изменение значения для name (собственный метод)
	u.name = name1
}

func (u User) isElder() bool { // Проверка на возраст
	a := u.age
	isTrue := false
	if a >= 18 {
		isTrue = true
	} else if a < 18 {
		isTrue = false
	}
	return isTrue
}

func (u User) getHighScore() int { // Наибольшее число очков
	high := 0
	for _, sc := range u.score {
		if high < sc {
			high = sc
		}
	}
	return high
}

func main() {
	// Структура
	var user User = User{name: "Max", age: 27, password: "1234", score: []int{23, 67, 89, 102, 345, 1}}
	// Проще так:
	// user := User{"Max", 27, "1234", []int{23, 67, 89, 102, 345, 1}}
	fmt.Println(user)
	fmt.Println(user.age)
	user.age = 40
	fmt.Println(user.age)
	user.setName("Den")
	fmt.Println(user.name)           // Изменение имени "Max" на "Den"
	fmt.Println(user.isElder())      // Проверка на возраст
	fmt.Println(user.getHighScore()) // Наибольшее число очков
}
