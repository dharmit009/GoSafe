package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// create a new Fyne app
	fmt.Println("HELLO")
	a := app.New()

	// create a new window
	w := a.NewWindow("My Window")

	// create a label widget
	label := widget.NewLabel("Hello, World!")

	// set the content of the window to the label widget
	w.SetContent(label)

	// show the window
	w.ShowAndRun()
}
