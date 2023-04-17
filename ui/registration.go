package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	rpass1 = widget.NewPasswordEntry()
	rpass2 = widget.NewPasswordEntry()
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Password Manager")
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.SetFixedSize(true)

	rpass1.SetPlaceHolder("Enter Master Password")
	rpass2.SetPlaceHolder("Confirm your Master Password")

	// Create the confirmation button
	confirmButton := widget.NewButton("Submit", createAccount(myWindow))

	// Create the rpass1 form layout
	form := container.New(layout.NewVBoxLayout(),
		rpass1,
		rpass2,
		confirmButton,
	)

	// Set the rpass1 form as the content of the window
	myWindow.SetContent(form)
	// Set the window size and icon myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.SetIcon(theme.FyneLogo())
	// Show the window and run the application
	myWindow.ShowAndRun()
}

func createAccount(myWindow fyne.Window) func() {
	return func() {
		if rpass1.Text == "" || rpass2.Text == "" {
			dialog.ShowInformation("Error", "Please enter a master password and confirm it.", myWindow)
			return
		}

		if rpass1.Text != rpass2.Text {
			dialog.ShowInformation("Error", "Master password and confirmation do not match.", myWindow)
			return
		}

		if rpass1.Text == rpass2.Text {
		dialog.ShowInformation("Success", "Account created successfully.", myWindow)
    }

	}
}
