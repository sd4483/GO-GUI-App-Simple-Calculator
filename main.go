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
	
)

/* func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03 :04 :05")
	clock.SetText(formatted)
} */

func main() {
	a := app.New()
	w := a.NewWindow("Hey!")
	
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text ...")

	content := binding.NewString()
	content.Set("")
	
	text := widget.NewLabelWithData(content)
	text.TextStyle = fyne.TextStyle{Bold: true}

	print_content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was: ", input.Text)
		content.Set(input.Text)
	}),text)
	
	w.SetContent(print_content)

	w.Resize(fyne.NewSize(300, 100))
	w.Show()
	a.Run()

/* 	clock := widget.NewLabel(" ")	
	updateTime(clock)
	
	w.SetContent(clock)

	go func() {
		for range time.Tick(time.Second){
			updateTime(clock)
		}
	}()
	
	w.SetMaster()
	w.Show()
	
	w2 := a.NewWindow("Larger")
	
	w2.SetContent(widget.NewButton("Open new window", func() {
		w3 := a.NewWindow("Second Window")
		w3.SetContent(widget.NewLabel("Second Window!"))
		w3.Show()
	}))
	
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()
	
	a.Run() */
}