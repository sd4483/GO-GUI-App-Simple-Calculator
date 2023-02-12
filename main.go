package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"log"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"strconv"
	
)

func conv_to_int (s string) int {
	i, e := strconv.Atoi(string(s))
	_ = e
	return i
}

/* func render_content () {
	log.Println("Content was: ", num_1.Text)
	log.Println("Content was: ", num_2.Text)
	content_1.Set("Entered number 01:   " + num_1.Text)
	content_2.Set("Entered number 02:   " + num_2.Text)

	n1 = conv_to_int(num_1.Text)
	n2 = conv_to_int(num_2.Text)
} */

func main() {
	a := app.New()
	w := a.NewWindow("Simple GUI Calculator")
	
	num_1 := widget.NewEntry()
	num_1.SetPlaceHolder("Enter number 01 ")

	num_2 := widget.NewEntry()
	num_2.SetPlaceHolder("Enter number 02 ")

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

	add_button := widget.NewButton("Add", func() {
		
		log.Println("Content was: ", num_1.Text)
		log.Println("Content was: ", num_2.Text)
		content_1.Set("Entered number 01:   " + num_1.Text)
		content_2.Set("Entered number 02:   " + num_2.Text)

		n1 := conv_to_int(num_1.Text)
		n2 := conv_to_int(num_2.Text)

		add := n1 + n2

		result.Set("Result of the addition:   " + strconv.Itoa(add))

	})

	sub_button := widget.NewButton("Subtract", func() {
		
		log.Println("Content was: ", num_1.Text)
		log.Println("Content was: ", num_2.Text)
		content_1.Set("Entered number 01:   " + num_1.Text)
		content_2.Set("Entered number 02:   " + num_2.Text)

		n1 := conv_to_int(num_1.Text)
		n2 := conv_to_int(num_2.Text)

		sub := n1 - n2

		result.Set("Result of the subtraction:   " + strconv.Itoa(sub))

	})

	mul_button := widget.NewButton("Multiply", func() {
		
		log.Println("Content was: ", num_1.Text)
		log.Println("Content was: ", num_2.Text)
		content_1.Set("Entered number 01:   " + num_1.Text)
		content_2.Set("Entered number 02:   " + num_2.Text)

		n1 := conv_to_int(num_1.Text)
		n2 := conv_to_int(num_2.Text)

		mul := n1 * n2

		result.Set("Result of the multiplication:   " + strconv.Itoa(mul))

	})

	div_button := widget.NewButton("Divide", func() {
		
		log.Println("Content was: ", num_1.Text)
		log.Println("Content was: ", num_2.Text)
		content_1.Set("Entered number 01:   " + num_1.Text)
		content_2.Set("Entered number 02:   " + num_2.Text)

		n1 := conv_to_int(num_1.Text)
		n2 := conv_to_int(num_2.Text)

		div := float64(n1) / float64(n2)

		result.Set("Result of the division:   " + strconv.FormatFloat(div, 'g', 3, 64))

	})

	layout_1 := container.NewVBox(num_1, num_2)
	layout_2 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), add_button, sub_button, mul_button, div_button, layout.NewSpacer())
	layout_3 := container.NewVBox(text_1, text_2, result_text)

	w.SetContent(container.New(layout.NewVBoxLayout(), layout_1, layout_2, layout_3))

	//w.Resize(fyne.NewSize(600, 200))
	w.Show()
	a.Run()
}