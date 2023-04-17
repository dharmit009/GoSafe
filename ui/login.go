package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Password Manager")
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.SetFixedSize(true)

	// Create the login form
	login := widget.NewEntry()
	login.SetPlaceHolder("Enter Master Password")
	login.Password = true

	// Create the login button
	loginButton := widget.NewButton("Login", func() {
		// Check if the master password entered is correct
		if login.Text == "test" {
			// Show the password manager interface
			// Replace this with your actual code to show the password manager
        myWindow.SetContent(container.New(layout.NewVBoxLayout(),
				widget.NewButton("Logout", func() {
					// Implement your logout functionality here
				}),
			))
		} else {
			// Show an error message if the master password is incorrect
			widget.NewLabel("Incorrect Master Password!")
		}
	})

	// Create the login form layout
	form := container.New(layout.NewVBoxLayout(),
		login,
		loginButton,
	)

	// Set the login form as the content of the window
	myWindow.SetContent(form)

	// Set the window size and icon
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.SetIcon(theme.FyneLogo())

	// Show the window and run the application
	myWindow.ShowAndRun()
}

