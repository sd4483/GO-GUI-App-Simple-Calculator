package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/canvas"
	"log"
	"fyne.io/fyne/v2/data/binding"
	//"image/color"
	//"time"
	//"strings"
	"strconv"
	
)

func main() {
	a := app.New()
	w := a.NewWindow("Hey!")
	
	num_1 := widget.NewEntry()
	num_1.SetPlaceHolder("Enter text ...")

	num_2 := widget.NewEntry()
	num_2.SetPlaceHolder("Enter text...")

	content_1 := binding.NewString()
	content_1.Set("")
	
	content_2 := binding.NewString()
    content_2.Set("")

	result := binding.NewString()
	result.Set("")
	
	text_1 := widget.NewLabelWithData(content_1)
	text_2 := widget.NewLabelWithData(content_2)

	result_text := widget.NewLabelWithData(result)
	result_text.TextStyle = fyne.TextStyle{Bold: true}

	print_content := container.NewVBox(num_1, num_2, widget.NewButton("Save", func() {
		log.Println("Content was: ", num_1.Text)
		log.Println("Content was: ", num_2.Text)
		content_1.Set("Entered number 01:   " + num_1.Text)
		content_2.Set("Entered number 02:   " + num_2.Text)

		i1, e1 := strconv.Atoi(string(num_1.Text))
		i2, e2 := strconv.Atoi(string(num_2.Text))
		add := i1 + i2
		_ = e1
		_ = e2

		result.Set("Result of the operation:   " + strconv.Itoa(add))

	}), text_1,text_2,result_text)
	
	w.SetContent(print_content)

	w.Resize(fyne.NewSize(600, 200))
	w.Show()
	a.Run()
}