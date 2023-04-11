package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/dharmit009/gopass/passutil"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My Password Manager")

	// Create a top canvas with an empty form
	topCanvas := container.NewHBox()

	form := widget.NewForm(
		&widget.FormItem{Text: "Website Name", Widget: widget.NewEntry()},
		&widget.FormItem{Text: "Username", Widget: widget.NewEntry()},
		&widget.FormItem{Text: "Password", Widget: widget.NewPasswordEntry()},
	)

	form.Append("Generate Password", widget.NewButton("Generate", func() {
		password := passutil.GeneratePassword() // generate a 16 character password using passutil package
		form.Items[2].Widget.(*widget.Entry).SetText(password)
	}))

	form.Hide() // hide the form initially

	// Create a bottom container with four buttons
	addBtn := widget.NewButton("Add", func() {
		// show the form when the button is clicked
		form.Show()
	})
	viewBtn := widget.NewButton("View", func() {})
	delBtn := widget.NewButton("Delete", func() {})
	updateBtn := widget.NewButton("Update", func() {})

	bottomContainer := container.NewVBox(
		addBtn,
		viewBtn,
		delBtn,
		updateBtn,
	)
	bottomContainer.Resize(fyne.NewSize(0, 50)) // set the height of the bottom container

	// Combine the top and bottom canvases in a VBox
	content := container.NewHBox(
		bottomContainer,
		form,
		topCanvas,
	)

	// Set the window content to the VBox
	myWindow.SetContent(content)

	myWindow.Resize(fyne.NewSize(400, 400)) // set a default window size

	myWindow.ShowAndRun()
}
