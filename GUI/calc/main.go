package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Калькулятор")
	window.Resize(fyne.NewSize(270, 320))
	label1 := widget.NewLabel("Введите первое число")
	textbox1 := widget.NewEntry()
	label2 := widget.NewLabel("Введите второе число")
	textbox2 := widget.NewEntry()

	answer := widget.NewLabel("")

	btn := widget.NewButton("Расчитать", func() {
		num1, err1 := strconv.ParseFloat(textbox1.Text, 64)
		num2, err2 := strconv.ParseFloat(textbox2.Text, 64)

		if err1 != nil || err2 != nil {
			answer.SetText("Ошибка ввода")
		} else {
			sum := num1 + num2
			sub := num1 - num2
			mul := num1 * num2
			div := num1 / num2
			answer.SetText(fmt.Sprintf("(+) %f\n(-) %f\n(*) %f\n(/) %f", sum, sub, mul, div))
		}
	})

	window.SetContent(container.NewVBox(
		label1,
		textbox1,
		label2,
		textbox2,
		btn,
		answer,
	))
	window.ShowAndRun()
}
