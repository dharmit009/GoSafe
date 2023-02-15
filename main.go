package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("### GOPASS ###")
	myLabel := widget.NewLabel("Hello, World")
	myContainer := container.NewCenter(myLabel)
	myWindow.SetContent(myContainer)
	myWindow.ShowAndRun()
}
