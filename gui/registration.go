package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dharmit009/gopass/passutil"
)

var (
  fileName = "master.enc"
	registered = false
	rpass1     = widget.NewPasswordEntry()
	rpass2     = widget.NewPasswordEntry()
)

func Registration(window fyne.Window) (*fyne.Container, bool) {

	rpass1.SetPlaceHolder("Enter Master Password")
	rpass2.SetPlaceHolder("Confirm your Master Password")

	// Create the confirmation button
	confirmButton := widget.NewButton("Submit", func(){
		flag := passutil.CheckPassEqual(rpass1.Text, rpass2.Text)
		if !flag {
			rpass1.SetText("")
			rpass2.SetText("")
      ShowErrorDialog(window, "Error", "Passwords do not match!")
			return
		}
		if passutil.CreateMasterKey(fileName, rpass2.Text, rpass1.Text) {
      container, registered := Login(window)
      window.SetContent(container)
      _ = registered
		} else {
			rpass1.SetText("")
			rpass2.SetText("")
      ShowErrorDialog(window, "Error", "Passwords must be equal to or more than 8 characters")
		}
	})


	// Create the rpass1 form layout
	c := container.New(layout.NewVBoxLayout(),
		rpass1,
		rpass2,
		confirmButton,
	)

	return c, registered
}
