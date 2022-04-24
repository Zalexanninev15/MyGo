package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Title")
	label1 := widget.NewLabel("Hello World!")
	textbox := widget.NewEntry()
	btn := widget.NewButton("Click on Me!", func() {
		data := textbox.Text
		fmt.Println(data)
		label1.SetText(data)
	})
	window.SetContent(container.NewVBox(
		label1,
		btn,
		textbox,
	))
	window.ShowAndRun()
}
