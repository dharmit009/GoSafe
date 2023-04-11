package gui

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dharmit009/gopass/watchman"
)

const (
	windowTitle = "Gopass"
	fileName    = "./test.enc"
)

func RunGUI() {
	a := app.New()
	w := a.NewWindow(windowTitle)

	passwordEntry := widget.NewPasswordEntry()
	passwordConfirmEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Enter your password"
	passwordConfirmEntry.PlaceHolder = "Re-enter your password"
  passwordEntry.FocusGained()

	form := widget.NewForm()

	if _, err := os.Stat(fileName); os.IsNotExist(err) {

		form.Append("", passwordEntry)
		form.Append("", passwordConfirmEntry)
		form.Append("", widget.NewButton("Create Master Key", func() {
			flag := watchman.CheckPassEqual(passwordEntry.Text, passwordConfirmEntry.Text)
			if flag != true {
				passwordEntry.SetText("")
				passwordConfirmEntry.SetText("")
				return
			}

			if watchman.CreateMasterKey(fileName, passwordConfirmEntry.Text, passwordEntry.Text) == false {
				passwordEntry.SetText("")
				passwordConfirmEntry.SetText("Reswapn")
			} else {
				// form.Hide()
			}
		}))

	} else {
    passwordConfirmEntry := widget.NewPasswordEntry()
    passwordConfirmEntry.PlaceHolder = "Enter Your Master Password"
		form.Append("", passwordConfirmEntry)
		form.Append("", widget.NewButton("Login", func() {
      comparer := watchman.CheckPassEqualToMP(passwordConfirmEntry.Text)
      if comparer == false{
        w.SetContent(widget.NewLabel("Login Unsuccessful"))
      }else if comparer == true {
        w.SetContent(widget.NewLabel("Login Successful!"))
      }else{
        w.SetContent(widget.NewLabel("Login Error"))
      }
		}))
    
  }

	container := fyne.NewContainerWithLayout(layout.NewMaxLayout(), form)
	w.SetContent(container)
	w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()
}
