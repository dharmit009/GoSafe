package gui

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

  "github.com/dharmit009/gopass/passutil"
)

var (
	logged = false
)

func Login(window fyne.Window) (*fyne.Container, bool) {

	rpass2 := widget.NewPasswordEntry()
	rpass2.SetPlaceHolder("Enter Master Password")

	// Create the login button
	loginButton := widget.NewButton("Login", func() {
		if comparer := passutil.CheckPassEqualToMP(rpass2.Text); comparer == true {
			logged = true
      container := Tabs(window)
      window.SetContent(container)
		} else {
			ShowErrorDialog(window, "Error", "Incorrect Password. Try Again")
		}

	})

	// Create the login form layout
	c := container.New(layout.NewVBoxLayout(),
		rpass2,
		loginButton,
	)

	return c, logged

}
